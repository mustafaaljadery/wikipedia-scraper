# Wikipedia Scraper

The goal of this project is to scrape Wikipedia and format the data so that it is readable by LLMs (tokenize the data).

![GIF demo of Scraper and Tokenizer](/images/demo.gif)

## How does it work

I wrote a detailed paper [here](https://www.maxaljadery.com/posts/scrape-wikipedia). It's actually a simple web scraper system design.

There are some design choices that I chose, that depending on what I'm building or what I want to do with this, I would do differently.

Firstly, concurrent locks on queue. Right now, I'm not processing the scrapping concurrently, however the extraction and adding to the queue is done concurrently. I'm not scraping concurrently because I don't want to put a lot of load on the Wikipedia server as this is just a project for me, I'm not really going to do anything useful with the data.

Secondly, storing the data. Firstly, if I wanted to store the data, I would just store the tokens. You can think of the tokens as a compressions of all the text. You will get way less tokens then you will have words, and they take less memory. (Given my intent is to use them for LLM training)

## FAQ

**Why Go?** For these types of problems, Go is really good. The networking and concurrency libraries that Go offers is amazing.

**What is a tokenizer?** A tokenizer is a preprocessing step before feeding your data into the model. It takes the words that you want to input in the model and splits them into chuncks, then transforms them into numbers so that they can be understood by the transformer model.

## Usage

Change the `.env.example` to `.env` then input your redis url and the start wikipedia URL, make sure the first one is semi-popular!

Make sure you are in the scraper folder.

To run the scrape, comment the tokenizer in main and keep the scraper.

To run the tokienzer, comment the scraper in main and keep the tokenizer.

**Run code**

```bash copy
go run main.go
```
