package middleware

import (
	"BooksWebservice/data"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyBook struct{}

//middleware stuff below
//mb remove the receiver - *Books here, and call file books.middleware
func NewBookValidation(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var book data.Book
		err := book.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Book is invalid.", http.StatusBadRequest)
			return
		} else if book.Id != primitive.NilObjectID { //the _id has to be an empty string in the new book
			http.Error(w, "You can't assign book id.", http.StatusBadRequest)
			return
		}
		//pass book with context of the request
		ctx := context.WithValue(r.Context(), KeyBook{}, book)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

/*Middleware for PUT request on /books/{id} route. Validates Book object, makes sure that user didn't mention id anywhere but in the URL.
Puts the book object in the context of the request and sends it along.*/
func UpdateBookValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var book data.Book
		vars := mux.Vars(r)
		id := vars["id"]

		err := book.FromJSON(r.Body)
		if err != nil {
			println(err.Error())
			http.Error(w, "Book is invalid.", http.StatusBadRequest)
			return
		} else if book.Id != primitive.NilObjectID {
			http.Error(w, "Book's _id must be empty", http.StatusBadRequest)
			return
		}
		book.Id, _ = primitive.ObjectIDFromHex(id) //will never be an error here.

		//pass book with context of the request
		ctx := context.WithValue(r.Context(), KeyBook{}, book)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
