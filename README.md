# Wikipedia Scraper

The goal of this project is to scrape Wikipedia and format the data so that it is readable by LLMS (tokenize the data).

## How does it work

I wrote a detailed paper [here](https://www.maxaljadery.com/posts/scrape-wikipedia). It's actually a simple web scraper system design.

There are some design choices that I chose, that depending on what I'm building or what I want to do with this, I would do differently.

Firstly, concurrent locks on queue. Right now, I'm not processing the scrapping concurrently, however the extraction and adding to the queue is done concurrently. I'm not scraping concurrently because I don't want to put a lot of load on the Wikipedia server as this is just a project for me, I'm not going to build a project around this data.

## FAQ

**Why Go?** For these types of problems, Go is really good. The networking and concurrency libraries that Go offers is amazing.

**What is a tokenizer?** A tokenizer is a preprocessing step before feeding your data into the model. It takes the words that you want to input in the model and splits them into chuncks, then transforms them into numbers so that they can be understood by the transformer model.

## Usage

Change the `.env.example` to `.env` then input your redis url and the start wikipedia URL, make sure the first one is semi-popular!

**To run the scraper**

Make sure you are in the scraper folder.

```bash copy
go run main.go
```

**To run the tokenizer**

Make sure you have data in the scraper, and you are in the tokenizer folder.

```bash copy
go run main.go
```
