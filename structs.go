package main

// clip data structure from API
type clip struct {
	Gameclipid   string `json:"gameClipId"`
	Gameclipuris []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"gameClipUris"`
	Titlename string `json:"titleName"`
}

// formatted structure for a clip after relevant metadata is extracted
type formattedClip struct {
	ID        string
	GameTitle string
	URI       string
}

// screenshot data structure from API
type screenshot struct {
	Screenshotid   string `json:"screenshotId"`
	Screenshoturis []struct {
		URI      string  `json:"uri"`
		Filesize float64 `json:"fileSize"`
		Uritype  string  `json:"uriType"`
	} `json:"screenshotUris"`
	Titlename string `json:"titleName"`
}

// formatted structure for a screenshot after relevant metadata is extracted
type formattedScreenshot struct {
	ID        string
	GameTitle string
	URI       string
}
