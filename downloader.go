package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func orchestrateDownloads(clips []formattedClip, screenshots []formattedScreenshot, downloadPath string) string {
	wg := sync.WaitGroup{}
	concurrencyLimiter := make(chan int, 5)

	folderPath := prepareDir(downloadPath)
	totalCount := len(clips) + len(screenshots)
	printProgress(0, totalCount)

	for i, screenshot := range screenshots {
		concurrencyLimiter <- i + 1
		wg.Add(1)
		go download(screenshot.URI, fmt.Sprintf("%s/%s_%s.png", folderPath, screenshot.GameTitle, screenshot.ID), &wg, concurrencyLimiter, totalCount)
	}

	i := len(clips)
	for _, clip := range clips {
		concurrencyLimiter <- i + 1
		wg.Add(1)
		go download(clip.URI, fmt.Sprintf("%s/%s_%s.mp4", folderPath, clip.GameTitle, clip.ID), &wg, concurrencyLimiter, totalCount)
		i++
	}

	wg.Wait()
	fmt.Print("\n")
	return folderPath
}

func download(URI string, filePath string, wg *sync.WaitGroup, concurrencyLimiter chan int, totalCount int) {
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

	currentCount := <-concurrencyLimiter
	printProgress(currentCount, totalCount)
	wg.Done()
}
