package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

func getClips(xuid string, authToken string) []formattedClip {
	fmt.Println("Getting user clips...")

	res := authGetRequest("/"+url.QueryEscape(xuid)+"/game-clips", authToken)

	userClips := []clip{}
	err := json.NewDecoder(res.Body).Decode(&userClips)
	if err != nil {
		log.Fatal("Error decoding API response for clips.")
	}
	defer res.Body.Close()

	for res.Header.Get("X-Continuation-Token") != "" {
		res = authGetRequest("/"+url.QueryEscape(xuid)+"/game-clips?continuationToken="+res.Header.Get("X-Continuation-Token"), authToken)
		defer res.Body.Close()

		extraClips := []clip{}
		err := json.NewDecoder(res.Body).Decode(&extraClips)
		if err != nil {
			log.Fatal("Error decoding continued API response for clips.")
		}
		userClips = append(userClips, extraClips...)
	}

	return extractClipsMetadata(&userClips, "clips")
}

func extractClipsMetadata(userClips *[]clip, filetype string) []formattedClip {
	formattedClips := []formattedClip{}
	bytes := 0.0

	for _, clip := range *userClips {
		formattedClip := formattedClip{}
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

	fmt.Printf("Found %d clips that will use up %.2fGB of storage space.\n", len(formattedClips), bytes/1073741824)
	return formattedClips
}
