package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/googollee/go-socket.io"
	"github.com/jaytaylor/html2text"
	_ "github.com/mattn/go-sqlite3"
)

type Visit struct {
	Url        string
	VisitCount int
}

type Word struct {
	Word  string
	Count *int
}

type ByCount []Word

func (slice ByCount) Len() int {
	return len(slice)
}

func (slice ByCount) Less(i, j int) bool {
	return *slice[i].Count > *slice[j].Count
}

func (slice ByCount) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	server.On("connection", ProcessHistory)

	/*
		server.On("connection", func(so socketio.Socket) {
			fmt.Println("on connection")
			so.Join("chat")
			so.On("chat message", func(msg string) {
				log.Println("emit:", so.Emit("chat message", msg))
				so.BroadcastTo("chat", "chat message", msg)
			})
			so.On("disconnection", func() {
				log.Println("on disconnect")
			})
		})
	*/

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)

}

func ProcessHistory(so socketio.Socket) {

	var words = []Word{}

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

		//go ProcessUrl(visit.Url)
		visit_words := ProcessUrl(visit.Url)

		if visit_words != nil {
		Loop:
			for _, visit_word := range visit_words {
				for _, word := range words {
					if word.Word == visit_word[0] {
						*word.Count++
						//fmt.Println("Repetido:", word.Word, "Número:", *word.Count)
						continue Loop
					}
				}
				aux := new(int)
				*aux = 1

				words = append(words, Word{visit_word[0], aux})
				//fmt.Println("Novo:", visit_word[0])
			}
		}

		sort.Sort(ByCount(words))

		so.Emit("words", words)

		/*
			file, err := os.Create("result.txt")
			if err != nil {
				panic(err)
				return
			}
			defer file.Close()

			for _, word := range words {
				file.WriteString(word.Word + " : " + strconv.Itoa(*word.Count) + "\n")
			}
		*/
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

		page_words := regex.FindAllStringSubmatch(body_text, -1)

		return page_words
	}

	return nil
}
