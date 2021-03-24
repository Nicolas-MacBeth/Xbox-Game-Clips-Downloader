package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
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
