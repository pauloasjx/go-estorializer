package main

import (
	"database/sql"
	"fmt"
	"strconv"

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
		fmt.Println(visit.Url + " - " + strconv.Itoa(visit.VisitCount))
	}
}

func ProcessUrl() {

}
