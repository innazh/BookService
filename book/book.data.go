package book

import (
	"BooksWebservice/database"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const dbName = "Books"
const collectionName = "books"

//all gets
//my error here may be that I'm using the same context for Find() and All()
func getBookList() ([]Book, error) {
	var bookList []Book
	// //get collection handler
	cur, err := database.DbConn.Database(dbName).Collection(collectionName).Find(context.Background(), bson.D{})
	//get all books from the database
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer cur.Close(context.Background())

	err = cur.All(context.TODO(), &bookList)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return bookList, nil
}

func getBookById() {

}

func getBooksByGenre() {

}

//or maybe within a certain decade, would only accept numbers that are devisible by 10
func getBooksByYear() {

}

//maybe redirect to the page with the inserted book
func insertBook(b Book) (string, error) {
	if &b == nil {
		return "", errors.New("book is empty/invalid")
	}
	if b.Id != primitive.NilObjectID {
		return "", errors.New("the ID needs to be empty")
	}

	var insertedId string
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := database.DbConn.Database(dbName).Collection(collectionName).InsertOne(ctx, b)
	if err != nil {
		return "", err
	}
	insertedId = result.InsertedID.(primitive.ObjectID).String()
	return insertedId, err
}
