package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// query API for given user's screenshots
func getScreenshots(xuid string, authToken string) []formattedScreenshot {
	fmt.Println("Getting user screenshots...")

	res := authGetRequest("/"+url.QueryEscape(xuid)+"/screenshots", authToken)

	// unmarshal response onto array of screenshot structs
	userScreenshots := []screenshot{}
	err := json.NewDecoder(res.Body).Decode(&userScreenshots)
	if err != nil {
		log.Fatal("Error decoding API response for screenshots.")
	}
	defer res.Body.Close()

	// continue querying API if a continuation token is present
	for res.Header.Get("X-Continuation-Token") != "" {
		res = authGetRequest("/"+url.QueryEscape(xuid)+"/screenshots?continuationToken="+res.Header.Get("X-Continuation-Token"), authToken)
		defer res.Body.Close()

		// append extra screenshots to previously queried screenshots
		extraScreenshots := []screenshot{}
		err := json.NewDecoder(res.Body).Decode(&extraScreenshots)
		if err != nil {
			log.Fatal("Error decoding continued API response for screenshots.")
		}
		userScreenshots = append(userScreenshots, extraScreenshots...)
	}

	// extract important data for each screenshot and format struct
	return extractScreenshotsMetadata(&userScreenshots)
}

// extract download URI & game title from screenshot data and aggregate total filesize
func extractScreenshotsMetadata(userScreenshots *[]screenshot) []formattedScreenshot {
	formattedScreenshots := []formattedScreenshot{}
	bytes := 0.0

	// loop over screenshots and format with extracted data
	for _, screenshot := range *userScreenshots {
		formattedScreenshot := formattedScreenshot{}

		// find download URI
		for _, uri := range screenshot.Screenshoturis {
			if uri.Uritype == "Download" {
				formattedScreenshot.URI = uri.URI
				formattedScreenshot.GameTitle = screenshot.Titlename
				formattedScreenshot.ID = screenshot.Screenshotid
				bytes += uri.Filesize
				break
			}
		}

		formattedScreenshots = append(formattedScreenshots, formattedScreenshot)
	}

	// tell user the total storage space that the screenshots will use
	fmt.Printf("Found %d screenshots that will use up %.2fGB of storage space.\n", len(formattedScreenshots), bytes/1073741824)
	return formattedScreenshots
}
