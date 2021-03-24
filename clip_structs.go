package main

type clip struct {
	Gameclipuris []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"gameClipUris"`
	Titlename string `json:"titleName"`
}

type formattedClip struct {
	gameTitle string
	URI       string
}
