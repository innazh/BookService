package book

import (
	"BooksWebservice/database"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
	//NEEDS SOME SORT OF DECODING, LOOK INTO THAT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	defer cur.Close(context.Background())
	err = cur.All(context.TODO(), &bookList)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(bookList[0].Name)
	return bookList, nil
}

func getBookById() {

}

func getBooksByGenre() {

}

//or maybe within a certain decade, would only accept numbers that are devisible by 10
func getBooksByYear() {

}
