package tokenizer

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tiktoken-go/tokenizer"
)


type Item struct {
	TagName string `json:"TagName"`
	Data    string `json:"Data"`
}

func GetDocs()[]string{
	var files_list []string
	f, _ := os.Open("./data")
	files, _ := f.Readdir(0)

	for _, file := range files {
		if !file.IsDir(){
			files_list = append(files_list,"./data/" +  file.Name())
		}
	}

	return files_list 
}

func GetDoc(path string)string{
	file, _ := os.Open(path)
	var content []byte
	defer file.Close()

	decoder := json.NewDecoder(file)

	var items []Item

	// Decode JSON data
	if err := decoder.Decode(&items); err != nil {
		log.Fatal(err)
	}


	for _, item := range items {
		if len(item.Data) > 10 {
			text := item.Data + "\n"
			content = append(content,[]byte(text)...)
		}
	}

	return string(content)
}

func Encode(str string)[]uint{
	enc, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		panic("tokenizer encode problem")
	}

	ids, _, _ := enc.Encode(str)

	return ids
}

func Decode(tokens []uint)string{
	enc, err := tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		panic("tokenizer decode problem")
	}

	text, _ := enc.Decode(tokens)

	return text
}