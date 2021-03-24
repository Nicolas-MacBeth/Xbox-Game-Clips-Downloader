package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

func getScreenshots(xuid string) []formattedScreenshot {
	fmt.Println("Getting user screenshots...")

	res := httpGetRequest("/" + url.QueryEscape(xuid) + "/screenshots")

	userScreenshots := []screenshot{}
	err := json.NewDecoder(res.Body).Decode(&userScreenshots)
	if err != nil {
		log.Fatal("Error decoding API response for screenshots.")
	}
	defer res.Body.Close()

	for res.Header.Get("X-Continuation-Token") != "" {
		res = httpGetRequest("/" + url.QueryEscape(xuid) + "/screenshots?continuationToken=" + res.Header.Get("X-Continuation-Token"))
		defer res.Body.Close()

		extraScreenshots := []screenshot{}
		err := json.NewDecoder(res.Body).Decode(&extraScreenshots)
		if err != nil {
			log.Fatal("Error decoding continued API response for screenshots.")
		}
		userScreenshots = append(userScreenshots, extraScreenshots...)
	}

	return extractScreenshotsMetadata(&userScreenshots)
}

func extractScreenshotsMetadata(userScreenshots *[]screenshot) []formattedScreenshot {
	formattedScreenshots := []formattedScreenshot{}
	bytes := 0.0

	for _, screenshot := range *userScreenshots {
		formattedScreenshot := formattedScreenshot{}
		for _, uri := range screenshot.Screenshoturis {
			if uri.Uritype == "Download" {
				formattedScreenshot.URI = uri.URI
				formattedScreenshot.gameTitle = screenshot.Titlename
				bytes += uri.Filesize
				break
			}
		}
		formattedScreenshots = append(formattedScreenshots, formattedScreenshot)
	}

	fmt.Printf("Found %d screenshots that will use up %.2fGB of storage space.\n", len(formattedScreenshots), bytes/1073741824)
	return formattedScreenshots
}
