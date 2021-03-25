package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func orchestrateDownloads(clips []formattedClip, screenshots []formattedScreenshot, downloadPath string) string {
	folderPath := prepareDir(downloadPath)
	totalCount := len(clips) + len(screenshots)
	printProgress(0, totalCount)

	for i, screenshot := range screenshots {
		download(screenshot.URI, fmt.Sprintf("%s/%s_%s.png", folderPath, screenshot.GameTitle, screenshot.ID))
		printProgress(i+1, totalCount)
	}

	i := len(clips)
	for _, clip := range clips {
		download(clip.URI, fmt.Sprintf("%s/%s_%s.mp4", folderPath, clip.GameTitle, clip.ID))
		printProgress(i+1, totalCount)
		i++
	}

	fmt.Print("\n")
	return folderPath
}

func download(URI string, filePath string) {
	res := downloadGetRequest(URI)
	defer res.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Could not create target file for download. Do you have the correct permissions?")
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal("Could not download file.")
	}
}
