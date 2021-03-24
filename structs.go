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

type screenshot struct {
	Screenshoturis []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"screenshotUris"`
	Titlename string `json:"titleName"`
}

type formattedScreenshot struct {
	gameTitle string
	URI       string
}
