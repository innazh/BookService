package handlers

import (
	"BooksWebservice/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Handles DELETE request on /books/{id} route. Deletes the book with {id} from the database.
Sends StatusOK header back to the client along with the id of deleted object*/
func (e *Env) DeleteBook(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle DELETE Book")
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

	if err = data.DeleteBook(objId, e.bookCollection); err != nil {
		e.l.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e.l.Println("Book with " + id + " has been deleted.")

	//case: the resource was deleted, everything is well
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
	return
}
