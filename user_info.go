package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	baseURL   = "https://xapi.us/v2"
	authToken = ""
)

func askUserForGamertag() string {
	fmt.Println("Please enter your Xbox network gamertag:")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gamertag := scanner.Text()
	return gamertag
}

func getXUID(gamertag string) string {
	fmt.Println("Getting XUID from gamertag...")

	res := httpGetRequest("/xuid/" + url.QueryEscape(gamertag))
	str, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error decoding API response.")
	}

	xuid := string(str)
	fmt.Printf("XUID found! (%s)\n", xuid)
	return xuid
}

func getClips(xuid string) {
	fmt.Println("Getting user clips...")

	res := httpGetRequest("/" + url.QueryEscape(xuid) + "/game-clips") //alt?

	userClips := []clip{}
	err := json.NewDecoder(res.Body).Decode(&userClips)
	if err != nil {
		log.Fatal("Error decoding API response.")
	}
	defer res.Body.Close()

	for res.Header.Get("X-Continuation-Token") != "" {
		res = httpGetRequest("/" + url.QueryEscape(xuid) + "/game-clips?continuationToken=" + res.Header.Get("X-Continuation-Token")) //alt?
		defer res.Body.Close()

		extraClips := []clip{}
		err := json.NewDecoder(res.Body).Decode(&extraClips)
		if err != nil {
			log.Fatal("Error decoding continued API response.")
		}
		userClips = append(userClips, extraClips...)
	}

	gatherClipsMetadata(&userClips)
}

func gatherClipsMetadata(userClips *[]clip) []formattedClip {
	formattedClips := []formattedClip{}
	bytes := 0.0

	for _, clip := range *userClips {
		formattedClip := formattedClip{}
		for _, uri := range clip.Gameclipuris {
			if uri.Uritype == "Download" {
				formattedClip.URI = uri.URI
				formattedClip.gameTitle = clip.Titlename
				bytes += uri.Filesize
				break
			}
		}
		formattedClips = append(formattedClips, formattedClip)
	}

	fmt.Printf("Found %d clips that will take %.2fGB of storage space.\n", len(formattedClips), bytes/1073741824)
	return formattedClips
}

func httpGetRequest(urlSuffix string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+urlSuffix, nil)
	req.Header.Set("X-AUTH", authToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP Get request error, did you make a typo in the gamertag?")
	}
	return res
}
