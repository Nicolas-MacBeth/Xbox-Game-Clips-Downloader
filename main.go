package main

import (
	"fmt"
)

func main() {
	greetUser()
	gamertag := askUserForGamertag()
	xuid := getXUID(gamertag)
	getClips(xuid)
}

func greetUser() {
	fmt.Println("Welcome to Nic MacBeth's Xbox network game clips downloader!")
	fmt.Println("Let's get started.")
}
