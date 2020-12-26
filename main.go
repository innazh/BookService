package main

import (
	"BooksWebservice/book"
	"BooksWebservice/database"
	"BooksWebservice/user"
	"BooksWebservice/utils"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	config := utils.GetConfiguration()

	database.SetupDB(config.ConnectionString) //connects to the database
	book.SetupRoutes(basePath)
	user.SetupRoutes(basePath)

	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseDbConn()
}
