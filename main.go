package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	word := os.Args[1]
	args := os.Args[2:]
	parseArgs(args)
	res, err := getWord(word)

	if err != nil {
		panic(err)
	}
	var m Result
	errs := json.Unmarshal([]byte(res), &m)
	if errs != nil {
		panic(err)
	}
	fmt.Println(m[0].Meanings[0])

}

func parseArgs(args []string) {
	for _, val := range args {
		fmt.Println(val)
	}
}

func getWord(word string) (sb string, err error) {
	resp, err := http.Get(dictionaryapi + word)

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
