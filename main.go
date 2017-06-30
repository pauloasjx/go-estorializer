package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"regexp"
	"sort"

	//"github.com/gorilla/mux"
	"github.com/jaytaylor/html2text"
)

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

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/test", ProcessHistory)

	http.ListenAndServe(":3000", nil)

}

func ProcessHistory(w http.ResponseWriter, req *http.Request) {

	//if req.Method == http.MethodPost {

	var words = []Word{}

	visit_words := ProcessUrl("http://google.com.br")

	if visit_words != nil {
	Loop:
		for _, visit_word := range visit_words {
			for _, word := range words {
				if word.Word == visit_word[0] {
					*word.Count++
					continue Loop
				}
			}
			aux := new(int)
			*aux = 1

			words = append(words, Word{visit_word[0], aux})
		}
	}

	sort.Sort(ByCount(words))

	http.Redirect(w, req, "/", http.StatusSeeOther)
	//}
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

		regex, err := regexp.Compile(`(\b[\p{L}A-Za-z]{4,16}\b)`)

		if err != nil {
			panic(err)
		}

		page_words := regex.FindAllStringSubmatch(body_text, -1)

		return page_words
	}

	return nil
}
