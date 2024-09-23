package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("no website, maxConcurrency or maxPages provided")
		return
	}
	rawBaseURL := os.Args[1]
	concurrency := os.Args[2]
	mPages := os.Args[3]
	maxConcurrency, err := strconv.Atoi(concurrency)
	if err != nil {
		fmt.Println("concurrency must be a number")
		return
	}
	maxPages, err := strconv.Atoi(mPages)
	if err != nil {
		fmt.Println("pages must be a number")
		return
	}
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}
