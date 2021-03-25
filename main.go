package main

import (
	"fmt"
	"strings"
)

// entrypoint for the entire project
func main() {
	greetUser()

	// gather download path, auth token and gamertag from the user
	downloadPath := strings.TrimSpace(askUserForInfo("Please enter the path where you'd like the download folder be created (relative or absolute).\nIf you enter nothing, the folder will be created in the current working directory."))
	authToken := askUserForInfo("Please visit https://xapi.us/ and enter your API key:")
	gamertag := askUserForInfo("Please enter your Xbox network gamertag:")

	// query API to get XUID, clips and screenshots
	xuid := getXUID(gamertag, authToken)
	clips := getClips(xuid, authToken)
	screenshots := getScreenshots(xuid, authToken)

	// download clips and screenshots
	dir := orchestrateDownloads(clips, screenshots, downloadPath)
	farewellUser(dir)
}

// give a proper welcome to the user :)
func greetUser() {
	fmt.Println("Welcome to Nic MacBeth's 'Xbox network' game clips downloader!")
	fmt.Println("Let's get started.")
}

// say goodbye to the beloved user :(
func farewellUser(path string) {
	fmt.Println("Download complete!")
	fmt.Printf("Your downloaded clips and screenshots can be found at %s\n", path)
	fmt.Println("Thanks for using this tool to download all your game DVR content!")
}
