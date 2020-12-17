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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const booksPath = "books"

func SetupRoutes(apiBasePath string) {
	booksHandlerO := http.HandlerFunc(booksHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, booksPath), booksHandlerO)                  //api/books
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, booksPath), http.HandlerFunc(bookHandler)) //api/books/
}

/*A handler for api/books route.
Possible methods: GET, POST
Function returns header Status not allowed for all other methods
case GET method: retrieves a slice of Books from the database, converts it to JSON format, sets the response header to application/json and writes out the list of books in the database
case POST method: retrieves the JSON object from the body of the response, decodes it, inserts it into the database, writes out the Id of the created object to the response*/
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
			log.Println("operation ReadAll: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &newb)
		if err != nil {
			log.Println("operation Unmarshal: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var insertedId string
		insertedId, err = insertBook(&newb)
		if err != nil {
			log.Println("operation insertBook: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if insertedId == "" {
			log.Println("Something went wrong, the Id of the inserted object is missing.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(insertedId))
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
		log.Println("operation getBookById: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if b == nil {
		log.Println("Couldn't locate the book with id =", bookId)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	//send the book to the user
	case http.MethodGet:
		var bookJSON []byte
		bookJSON, err = json.Marshal(b)
		if err != nil {
			log.Println("operation Marshal: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookJSON)
	case http.MethodPut:
		var updatedBook Book
		var bodyBytes []byte
		bodyBytes, err = ioutil.ReadAll(r.Body)

		err = json.Unmarshal(bodyBytes, &updatedBook)
		if err != nil {
			log.Println("operation Unmarshal: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = updateBook(bookId, updatedBook)
		if err != nil {
			log.Println("operation updateBook: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		//maybe redirect to the updated product page.
	case http.MethodDelete:
		err = deleteBook(bookId)
		if err != nil {
			log.Println("operation deleteBook: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
