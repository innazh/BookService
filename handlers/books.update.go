package handlers

import (
	"BooksWebservice/data"
	"BooksWebservice/middleware"
	"net/http"
)

/*Handles PUT request to /books/{id}. TODO: finish writing this comment*/
//for this function we'll use middlerware to validate the book.
func (e *Env) UpdateBook(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle PUT Book")
	var err error
	var book data.Book

	book = r.Context().Value(middleware.KeyBook{}).(data.Book) //we validate book and ids using middleware

	if err = data.UpdateBook(book, e.bookCollection); err != nil {
		e.l.Println("UpdateBook: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
