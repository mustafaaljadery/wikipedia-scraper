package main

import (
	"fmt"
	"scraper/linkqueue"
	"scraper/scrape"
	"time"
)

func main(){
	q := linkqueue.Queue{}
	startTime := time.Now()
	q.LoadQueue()
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("Loaded data in: %v\n", duration)

	var startUrl string = "https://en.wikipedia.org/wiki/Web_scraping" 
	scrape.RunScrape(q, &startUrl)
}
