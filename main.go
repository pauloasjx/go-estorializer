package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
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

		ProcessUrl(visit.Url)
		return
	}
}

func ProcessUrl(url string) {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

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

	body_text, err := html2text.FromString(string(body))

	fmt.Println(body_text)
}
