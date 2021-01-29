package handlers

import (
	"BooksWebservice/data"
	"BooksWebservice/middleware"
	"net/http"
)

/*Handles POST request on /books route. Retrieves book from context of the request and adds it to the database.
Sends StatusCreated header back to the client along with book Id*/
func (e *Env) AddBook(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle POST Book")
	var err error
	book := r.Context().Value(middleware.KeyBook{}).(data.Book) //object is validated via middleware --> NewBookValidation

	var insertedId string
	insertedId, err = data.InsertBook(&book, e.bookCollection)
	if err != nil {
		e.l.Println("failed to insert the book: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if insertedId == "" { //will probably never happen
		e.l.Println("Something went wrong, the Id of the inserted object is missing.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(insertedId))
	return
}
