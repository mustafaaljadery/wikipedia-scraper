package main

import (
	"fmt"
	"scraper/linkqueue"
	"scraper/scrape"
	"scraper/tokenizer"
	"time"
)

func Scrape(){
	q := linkqueue.Queue{}
	startTime := time.Now()
	q.LoadQueue(100)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Loaded data in: %v\n", duration)

	var startUrl string = "https://en.wikipedia.org/wiki/Web_scraping" 
	scrape.RunScrape(&q, &startUrl, 3)
}

func Tokenize(){
	docs := tokenizer.GetDocs()
	doc := tokenizer.GetDoc(docs[0])

	fmt.Println(tokenizer.Encode(doc))
}

func main(){
	//Scrape()
	Tokenize()
}