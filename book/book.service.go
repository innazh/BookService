package book

/*Some common response codes include:

400 Bad Request — Client sent an invalid request — such as lacking required request body or parameter
401 Unauthorized — Client failed to authenticate with the server
403 Forbidden — Client authenticated but does not have permission to access the requested resource
404 Not Found — The requested resource does not exist
412 Precondition Failed — One or more conditions in the request header fields evaluated to false
500 Internal Server Error — A generic error occurred on the server
503 Service Unavailable — The requested service is not available*/

/*Web API implements protocol specification and thus it incorporates concepts like caching, URIs, versioning, request/response headers, and various content formats in it.*/

/*TODO: make function check ERR, to make the code cleaner
maybe also functions that convert objectID to string and string to objectID
in database.go, make function that returns a collection
READ ABOUT CONTEXTS AND WHY WE NEED THEM, WHAT CAN THEY BE USED FOR*/
import (
	"BooksWebservice/cors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const booksPath = "books"

func SetupRoutes(apiBasePath string) {
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, booksPath), cors.ValidateMiddleware(booksHandler)) //api/books
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, booksPath), cors.ValidateMiddleware(bookHandler)) //api/books/
}

/*A handler for api/books route.
Possible methods: GET, POST
Function returns header Status not allowed for all other methods
case GET method: retrieves a slice of Books from the database, converts it to JSON format, sets the response header to application/json and writes out the list of books in the database
case POST method: retrieves the JSON object from the body of the response, decodes it, inserts it into the database, writes out the Id of the created object to the response*/
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//get all books
	case http.MethodGet:
		bookList, err := getBookList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		JSONbookList, err := json.Marshal(bookList)
		if err != nil {
			log.Println("error encoding bookList into JSON: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(JSONbookList)
	//create book
	case http.MethodPost:
		var newb Book
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("reading request body: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &newb)
		if err != nil || newb.Id != primitive.NilObjectID {
			log.Println("unmarshaling the body object, checking object's validity: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid object, please check your formatting -"))
			return
		}

		var insertedId string
		insertedId, err = insertBook(&newb)
		if err != nil {
			log.Println("failed to insert the book: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if insertedId == "" { //will probably never happen
			log.Println("Something went wrong, the Id of the inserted object is missing.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Inserted book's ID is " + insertedId))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	var b *Book
	var err error
	//get book from db
	bookId := strings.Split(r.URL.Path, "books/")[1]
	b, err = getBookById(bookId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if b == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Couldn't locate the book with id = " + bookId))
		return
	}

	switch r.Method {
	//send book to the user
	case http.MethodGet:
		var bookJSON []byte
		if bookJSON, err = json.Marshal(b); err != nil {
			log.Println("converting book to json: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookJSON)
	//modify book
	case http.MethodPut:
		var updatedBook Book
		var bodyBytes []byte
		if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
			log.Println("reading request body: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(bodyBytes, &updatedBook); err != nil {
			log.Println("unmarshaling the body" + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid object, please check your formatting"))
			return
		}

		if err = updateBook(bookId, updatedBook); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	//delete book
	case http.MethodDelete:

		if err = deleteBook(bookId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
