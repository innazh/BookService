package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const booksPath = "books"

func SetupRoutes(apiBasePath string) {
	booksHandlerO := http.HandlerFunc(booksHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, booksPath), booksHandlerO) //api/books
}

//all gets
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bookList, err := getBookList()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		JSONbookList, err := json.Marshal(bookList)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(JSONbookList)

	case http.MethodPost:
		var newb Book
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("after ReadAll:" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &newb)
		if err != nil {
			log.Println("after Unmarshal" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = insertBook(newb)
		if err != nil {
			log.Println("after insert" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {

}
