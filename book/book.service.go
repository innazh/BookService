package book

import (
	"fmt"
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
		_, err := getBookList()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {

}
