package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	//"sync"
	//"os"
	//"strconv"

	"github.com/jaytaylor/html2text"
	_ "github.com/mattn/go-sqlite3"
)

type Visit struct {
	Url        string
	VisitCount int
}

func main() {
	var words []string

	database, err := sql.Open("sqlite3", "./History")
	if err != nil {
		panic(err)
	}

	visits, err := database.Query("SELECT url, visit_count FROM urls")
	if err != nil {
		panic(err)
	}
	defer visits.Close()

	var visit Visit

	for visits.Next() {
		visits.Scan(&visit.Url, &visit.VisitCount)
		//fmt.Println(visit.Url + " - " + strconv.Itoa(visit.VisitCount))

		//go ProcessUrl(visit.Url)
		visit_words := ProcessUrl(visit.Url)

		if visit_words != nil {
			for _, visit_word := range visit_words {
				words = append(words, string(visit_word[0]))
			}
		}

		fmt.Println(words)
		//return

	}
}

func ProcessUrl(url string) [][]string {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {

		/*
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				panic(err)
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				panic(err)
			}
		*/

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		body_text, err := html2text.FromString(string(body))
		if err != nil {
			panic(err)
		}

		//var text string
		//json.Unmarshal([]byte(body_text), &text)

		//fmt.Println(text)

		regex, err := regexp.Compile(`(\b[\p{L}A-Za-z]{4,16}\b)`)
		//regex, err := regexp.Compile(`(\b[\p{L}A-Za-z]{4,}\b)`)
		//regex, err := regexp.Compile(`(\b[\p{L}A-Za-z]+\b)`)
		//regex, err := regexp.Compile(`(\b[^\s]+\b)`)
		if err != nil {
			panic(err)
		}

		words := regex.FindAllStringSubmatch(body_text, -1)

		return words
	}

	return nil
}
