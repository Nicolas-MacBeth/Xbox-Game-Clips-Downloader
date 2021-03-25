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
	baseURL      = "https://xapi.us/v2"   // base API URL
	downloadPath = "/xbox_DVR_downloads_" // folder name for downloads
)

// get request helper function for API queries that include an auth token
func authGetRequest(urlSuffix string, authToken string) *http.Response {
	// set provided token in headers
	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+urlSuffix, nil)
	req.Header.Set("X-AUTH", authToken)

	// send request and return response
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP Get request error, did you make a typo in the gamertag?")
	}
	return res
}

// get request helper without auth token, used for downloading
func downloadGetRequest(URI string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URI, nil)

	// send request and return response
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP Get request error, could not download data.")
	}
	return res
}

// create directory for downloads
func prepareDir(userPath string) string {
	// if user entered no path, add period to tell Go to create folder in current working directory
	if userPath == "" {
		userPath = "."
	}

	// path is built with the user-picked directory, the base folder name and the current date/time
	t := time.Now()
	folderPath := userPath + downloadPath + fmt.Sprintf("%d-%d-%d_%dh%dm%ds", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// test absolute path
	absPath, err := filepath.Abs(folderPath)
	if err != nil {
		log.Fatal("Unable to decode folder path on this system.")
	}

	// create target download directory
	err = os.Mkdir(absPath, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create target download folder. Did you make a typo, or are you lacking the correct permissions?")
	}

	return absPath
}

// print the current download progress, using \r to replace the current progress
func printProgress(current int, total int) {
	fmt.Printf("\r%d of %d items downloaded.", current, total)
}
