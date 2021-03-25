package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
)

// loop over clips & screenshots and download them in goroutines
func orchestrateDownloads(clips []formattedClip, screenshots []formattedScreenshot, downloadPath string) string {
	// limit amount of concurrent downloads to 5 using channel and use WaitGroup to keep program alive when downloading
	wg := sync.WaitGroup{}
	concurrencyLimiter := make(chan int, 5)
	invalidChars := regexp.MustCompile("[~\"#%&*:<>?/\\{|}]+")

	// check target download directory
	folderPath := prepareDir(downloadPath)
	totalCount := len(clips) + len(screenshots)
	printProgress(0, totalCount)

	// loop over screenshots and download each
	for i, screenshot := range screenshots {
		concurrencyLimiter <- i + 1
		wg.Add(1)
		finalPath := fmt.Sprintf("%s/%s_%s.png", folderPath, invalidChars.ReplaceAllString(screenshot.GameTitle, "_"), invalidChars.ReplaceAllString(screenshot.ID, "_"))
		go download(screenshot.URI, finalPath, &wg, concurrencyLimiter, totalCount)
	}

	// loop over clips and download each
	i := len(clips)
	for _, clip := range clips {
		concurrencyLimiter <- i + 1
		wg.Add(1)
		finalPath := fmt.Sprintf("%s/%s_%s.mp4", folderPath, invalidChars.ReplaceAllString(clip.GameTitle, "_"), invalidChars.ReplaceAllString(clip.ID, "_"))
		go download(clip.URI, finalPath, &wg, concurrencyLimiter, totalCount)
		i++
	}

	// wait for all downloads in goroutines to finish
	wg.Wait()
	fmt.Print("\n")
	return folderPath
}

// download a given clip or screenshot
func download(URI string, filePath string, wg *sync.WaitGroup, concurrencyLimiter chan int, totalCount int) {
	// send out get request
	res := downloadGetRequest(URI)
	defer res.Body.Close()

	// create file with proper path
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Could not create target file for download. Do you have the correct permissions?")
	}
	defer out.Close()

	// copy data into file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal("Could not download file.")
	}

	// use channel to keep track of progress and release WaitGroup item
	currentCount := <-concurrencyLimiter
	printProgress(currentCount, totalCount)
	wg.Done()
}
