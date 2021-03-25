package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	baseURL      = "https://xapi.us/v2"
	downloadPath = "/xbox_DVR_downloads_"
)

func authGetRequest(urlSuffix string, authToken string) *http.Response {
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

func prepareDir(userPath string) string {
	if userPath == "" {
		userPath = "."
	}

	t := time.Now()
	folderPath := userPath + downloadPath + fmt.Sprintf("%d-%d-%d_%dh%dm%ds", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	absPath, err := filepath.Abs(folderPath)
	if err != nil {
		log.Fatal("Unable to decode folder path on this system.")
	}

	err = os.Mkdir(absPath, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create target download folder. Did you make a typo, or are you lacking the correct permissions?")
	}

	return absPath
}

func printProgress(current int, total int) {
	fmt.Printf("\r%d of %d items downloaded.", current, total)
}
