package main

import (
	"log"
	"net/http"
)

const (
	baseURL   = "https://xapi.us/v2"
	authToken = ""
)

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
