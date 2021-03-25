package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// query API for given user's clips
func getClips(xuid string, authToken string) []formattedClip {
	fmt.Println("Getting user clips...")

	res := authGetRequest("/"+url.QueryEscape(xuid)+"/game-clips", authToken)

	// unmarshal response onto array of clip structs
	userClips := []clip{}
	err := json.NewDecoder(res.Body).Decode(&userClips)
	if err != nil {
		log.Fatal("Error decoding API response for clips.")
	}
	defer res.Body.Close()

	// continue querying API if a continuation token is present
	for res.Header.Get("X-Continuation-Token") != "" {
		res = authGetRequest("/"+url.QueryEscape(xuid)+"/game-clips?continuationToken="+res.Header.Get("X-Continuation-Token"), authToken)
		defer res.Body.Close()

		// append extra clips to previously queried clips
		extraClips := []clip{}
		err := json.NewDecoder(res.Body).Decode(&extraClips)
		if err != nil {
			log.Fatal("Error decoding continued API response for clips.")
		}
		userClips = append(userClips, extraClips...)
	}

	// extract important data for each clip and format struct
	return extractClipsMetadata(&userClips, "clips")
}

// extract download URI & game title from clip data and aggregate total filesize
func extractClipsMetadata(userClips *[]clip, filetype string) []formattedClip {
	formattedClips := []formattedClip{}
	bytes := 0.0

	// loop over clips and format with extracted data
	for _, clip := range *userClips {
		formattedClip := formattedClip{}

		// find download URI
		for _, uri := range clip.Gameclipuris {
			if uri.Uritype == "Download" {
				formattedClip.URI = uri.URI
				formattedClip.GameTitle = clip.Titlename
				formattedClip.ID = clip.Gameclipid
				bytes += uri.Filesize
				break
			}
		}

		formattedClips = append(formattedClips, formattedClip)
	}

	// tell user the total storage space that the clips will use
	fmt.Printf("Found %d clips that will use up %.2fGB of storage space.\n", len(formattedClips), bytes/1073741824)
	return formattedClips
}
