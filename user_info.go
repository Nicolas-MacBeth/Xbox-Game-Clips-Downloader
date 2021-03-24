package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

func askUserForInfo(requestedData string) string {
	fmt.Println(requestedData)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func getXUID(gamertag string, authToken string) string {
	fmt.Println("Getting XUID from gamertag...")

	res := authGetRequest("/xuid/"+url.QueryEscape(gamertag), authToken)
	str, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error decoding API response.")
	}

	xuid := string(str)
	fmt.Printf("XUID found! (%s)\n", xuid)
	return xuid
}
