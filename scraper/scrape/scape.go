package scrape

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"scraper/linkqueue"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type DataStruct struct {
    TagName  string
    Data string
}

func RunScrape(q *linkqueue.Queue,  link *string, concurrent_scrape int){
	var wg sync.WaitGroup

	for {
		wg.Add(1)
		go func() { 
			ScrapeData(q, link)
			wg.Done()
		 }()
		wg.Wait()
	}

}

func ScrapeData(q *linkqueue.Queue, link *string){
	url := ""
	if (link != nil && *link != ""){
		url = *link
		*link = ""
	} else if (q.HasNext()){
		url = q.Dequeue().(string)
	}

	fmt.Println("Scraping URL:", url)

	if url == "" {
		return 
	} 

	c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

	var wgs sync.WaitGroup

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		wgs.Add(2)

		go func(){ 
			ExtractData(q, e, url)
			wgs.Done()
		}()
		go func(){ 
			AddLinks(q, e)
			wgs.Done()
		}()
		wgs.Wait()
    })

	c.Visit(url)
}

func ExtractData(q *linkqueue.Queue, e *colly.HTMLElement, url string){
	data := []DataStruct{}

	e.ForEach("h1, h2, h3, h4, p, a", func(_ int, elem *colly.HTMLElement) {
		text := elem.Text
		if strings.HasSuffix(text, "[edit]") {
			text = strings.TrimSuffix(text, " [edit]")
		}
		data = append(data, DataStruct{TagName: elem.Name, Data: text})
	})

	encoded := base64.StdEncoding.EncodeToString([]byte(url))
	jsonData, _ := json.Marshal(data)
	file, _ := os.Create("./data/" + encoded +  ".json")
	file.Write(jsonData)
}

func AddLinks(q *linkqueue.Queue, e *colly.HTMLElement){
	links := e.ChildAttrs("a", "href")

	for index, value := range links {
		if (index == 5){
			break;
		} else if (strings.Contains(value, "wiki") && !strings.Contains(value, "https://") && !strings.Contains(value, ".svg") && !strings.Contains(value, ".jpg") && !strings.Contains(value, ".png")){
			q.Enqueue("https://en.wikipedia.org" + value)
		}
	}
}
