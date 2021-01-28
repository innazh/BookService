//All GET handlers for /Books
package handlers

import (
	"BooksWebservice/data"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Handles GET request on /books route. Sends back a collection of Books in JSON format.*/
//TODO use ToJSON function to convert books to JSON
func (e *Env) GetBooks(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle GET Books")
	w.Header().Set("Access-Control-Allow-Credentials", "true") //this fixed the credentials problem. Allows to accept requests with cookies(credentials)
	bookList, err := data.GetBooks(e.bookCollection)
	if err != nil {
		e.l.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = bookList.ToJSON(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e.l.Println("error encoding bookList into JSON: " + err.Error())
		return
	}
	return
}

/*Handles GET request on /books/{id} route. Sends back the book with {id}.*/
//TODO: rn if the book doesn't exist we send back null. Send back a msg 'resource doesn't exist'.
func (e *Env) GetBook(w http.ResponseWriter, r *http.Request) {
	var book *data.Book
	e.l.Println("Handle GET Book")
	vars := mux.Vars(r)
	id := vars["id"]

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Id."))
		return
	}
	book, err = data.GetBookById(objId, e.bookCollection)

	var bookJSON []byte
	if bookJSON, err = json.Marshal(book); err != nil {
		log.Println("converting book to json: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bookJSON)
	return
}
