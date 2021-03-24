package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	baseURL   = "https://xapi.us/v2"
	authToken = ""
)

func authGetRequest(urlSuffix string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+urlSuffix, nil)
	req.Header.Set("X-AUTH", authToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP Get request error, did you make a typo in the gamertag?")
	}
	return res
}

func downloadGetRequest(URI string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URI, nil)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP Get request error, could not download data.")
	}
	return res
}

func prepareDir() string {
	t := time.Now()
	folderPath := downloadPath + fmt.Sprintf("%d-%d-%d_%dh%dm%ds", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	err := os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Unable to create target download folder. Do you have the correct permissions?")
	}
	return folderPath
}

func printProgress(current int, total int) {
	fmt.Printf("\r%d of %d items downloaded.", current, total)
}
