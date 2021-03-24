package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	greetUser()
	authToken := askUserForInfo("Please visit https://xapi.us/ and enter your API key:")
	gamertag := askUserForInfo("Please enter your Xbox network gamertag:")
	xuid := getXUID(gamertag, authToken)
	clips := getClips(xuid, authToken)
	screenshots := getScreenshots(xuid, authToken)
	dir := orchestrateDownloads(clips, screenshots)
	farewellUser(dir)
}

func greetUser() {
	fmt.Println("Welcome to Nic MacBeth's 'Xbox network' game clips downloader!")
	fmt.Println("Let's get started.")
}

func farewellUser(dir string) {
	fmt.Println("Download complete!")
	path, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal("Error resolving path, but download was completed successfully.")
	}
	fmt.Printf("Your downloaded clips and screenshots can be found at %s\n", path)
	fmt.Println("Thanks for using this tool to download all your game DVR content!")
}
