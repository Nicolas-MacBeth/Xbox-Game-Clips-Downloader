package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

// get input from user with parametrized requested data description
func askUserForInfo(requestedData string) string {
	// print info request
	fmt.Println(requestedData)
	fmt.Print("> ")

	// get user input and return it
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

// fetch XUID from provided gamertag
func getXUID(gamertag string, authToken string) string {
	fmt.Println("Getting XUID from gamertag...")

	// send authenticated get request
	res := authGetRequest("/xuid/"+url.QueryEscape(gamertag), authToken)
	str, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error decoding API response.")
	}

	// return XUID
	xuid := string(str)
	fmt.Printf("XUID found! (%s)\n", xuid)
	return xuid
}
