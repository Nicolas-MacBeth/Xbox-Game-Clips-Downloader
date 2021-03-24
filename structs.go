package main

type clip struct {
	Gameclipid   string `json:"gameClipId"`
	Gameclipuris []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"gameClipUris"`
	Titlename string `json:"titleName"`
}

type formattedClip struct {
	ID        string
	GameTitle string
	URI       string
}

type screenshot struct {
	Screenshotid   string `json:"screenshotId"`
	Screenshoturis []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"screenshotUris"`
	Titlename string `json:"titleName"`
}

type formattedScreenshot struct {
	ID        string
	GameTitle string
	URI       string
}
