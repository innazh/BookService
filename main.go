package main

import (
	"BooksWebservice/authentication"
	"BooksWebservice/book"
	"BooksWebservice/database"
	"BooksWebservice/settings"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	settings := settings.GetSettings()

	database.SetupDB(settings.ConnectionString) //connects to the database
	book.SetupRoutes(basePath)
	authentication.SetupRoutes(basePath)

	err := http.ListenAndServe(settings.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseDbConn()
}
