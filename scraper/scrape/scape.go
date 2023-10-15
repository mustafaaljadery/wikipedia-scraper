package scrape

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"scraper/linkqueue"
	"strings"

	"github.com/gocolly/colly"
)

type DataStruct struct {
    TagName  string
    Data string
}

func RunScrape(q linkqueue.Queue,  link *string){
	fmt.Println("run")
	url := ""
	if (link != nil){
		url = *link
	} else{
		url = q.Dequeue().(string)
	}

	fmt.Println(url)
	c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
    )

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		ExtractData(q, e, url)
		AddLinks(q, e)
		var emptyValue *string
		RunScrape(q, emptyValue)
    })

	c.Visit(url)
}

func ExtractData(q linkqueue.Queue, e *colly.HTMLElement, url string){
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

func AddLinks(q linkqueue.Queue, e *colly.HTMLElement){
	links := e.ChildAttrs("a", "href")

	for _, value := range links {
		if (strings.Contains(value, "/wiki/") && !strings.Contains(value, "https://") && !strings.Contains(value, ".svg") && !strings.Contains(value, ".jpg") && !strings.Contains(value, ".png")){
			q.Enqueue("https://en.wikipedia.org" + value)
		}
	}
}
