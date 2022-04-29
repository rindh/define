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
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments!")
		os.Exit(0)
	}

	word := os.Args[1]
	fmt.Println(run(word))
}

func formatOutput(data Result) (res string) {
	res += data[0].Word + "\n"
	if len(data[0].Origin) > 0 {
		res += "Origin: " + data[0].Origin
	}

	res += "Meanings \n"
	for _, val := range data[0].Meanings {
		res += "Used as a " + val.PartOfSpeech + "\n"

		for _, k := range val.Definitions {
			res += "	Definition: " + k.Definition + "\n"
			if len(k.Example) > 0 {
				res += "	- Example: " + k.Example + "\n"
			}
		}
		res += "\n"
	}
	return
}

func run(word string) string {
	res, err := getWord(word)
	if err != nil {
		panic(err)
	}
	resultData, err := unmarshalRes(res)
	if err != nil {
		panic(err)
	}
	formatted := formatOutput(resultData)
	return formatted
}

func unmarshalRes(body string) (m Result, err error) {
	errs := json.Unmarshal([]byte(body), &m)
	if errs != nil {
		panic(err)
	}
	return
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
