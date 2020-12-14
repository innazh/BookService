package main

import (
	"BooksWebservice/book"
	"BooksWebservice/database"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	database.SetupDB() //connects to the database
	book.SetupRoutes(basePath)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
