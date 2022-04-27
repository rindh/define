package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result []struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text  string `json:"text"`
		Audio string `json:"audio,omitempty"`
	} `json:"phonetics"`
	Origin   string `json:"origin"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string        `json:"definition"`
			Example    string        `json:"example"`
			Synonyms   []interface{} `json:"synonyms"`
			Antonyms   []interface{} `json:"antonyms"`
		} `json:"definitions"`
	} `json:"meanings"`
}

var dictionaryapi = "https://api.dictionaryapi.dev/api/v2/entries/en/"

func main() {
	res, err := getWord()

	if err != nil {
		panic(err)
	}
	var m Result
	errs := json.Unmarshal([]byte(res), &m)
	if errs != nil {
		panic(err)
	}
	fmt.Println(m[0].Meanings)

}

func getWord() (sb string, err error) {
	println("Hello world!")
	resp, err := http.Get(dictionaryapi + "hello")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb = string(body)
	return
}
