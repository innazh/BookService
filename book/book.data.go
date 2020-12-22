package book

import (
	"BooksWebservice/database"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const dbName = "Books"
const collectionName = "books"

//all gets
func getBookList() ([]Book, error) {
	var bookList []Book
	bookColl := database.GetMongoDbCollection(dbName, collectionName)
	//get all books from the database
	cur, err := bookColl.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("find operation on the db collection: " + err.Error())
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var b Book
		err = cur.Decode(&b)
		bookList = append(bookList, b)
	}
	//check for the err after the loop
	if err != nil {
		log.Println("error while decoding the db results: " + err.Error())
		return nil, err
	}

	return bookList, nil
}

func getBookById(id string) (*Book, error) {
	// convert given ID to ObjectId
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("converting Id: " + err.Error())
		return nil, err
	}

	var b Book
	bookColl := database.GetMongoDbCollection(dbName, collectionName)
	filter := bson.M{"_id": objId}

	//check if the result isn't null (no matching documents were found)
	result := bookColl.FindOne(context.Background(), filter)
	if result == nil {
		return nil, nil
	}

	err = result.Decode(&b)
	if err != nil {
		log.Println("error while decoding the db results: ", err.Error())
		return nil, err
	}

	return &b, nil
}

func getBooksByGenre() {

}

//or maybe within a certain decade, would only accept numbers that are devisible by 10
func getBooksByYear() {

}

func insertBook(b *Book) (string, error) {
	var insertedId string
	bookColl := database.GetMongoDbCollection(dbName, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := bookColl.InsertOne(ctx, b)
	if err != nil {
		return "", err
	}
	insertedId = result.InsertedID.(primitive.ObjectID).Hex()
	return insertedId, err
}

func updateBook(id string, b Book) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("converting object Id: " + err.Error())
		return err
	}

	bookColl := database.GetMongoDbCollection(dbName, collectionName)
	filter := bson.M{"_id": objId} //filter := bson.M{"_id": bson.M{"$eq": objId}}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := bookColl.UpdateOne(ctx, filter, bson.D{{"$set", b}}) //because of omitempty, only the fields we fill in are updated
	if err != nil {
		log.Println("performing an update operation: " + err.Error())
		return err
	} else if result.ModifiedCount == 0 {
		msg := "Failed to update the Book with id = " + id
		log.Println(msg)
		return errors.New(msg)
	}
	return nil
}

func deleteBook(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("converting object Id: " + err.Error())
		return err
	}

	filter := bson.M{"_id": objId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	bookColl := database.GetMongoDbCollection(dbName, collectionName)
	result, err := bookColl.DeleteOne(ctx, filter)

	if err != nil {
		log.Println("performing a delete operation: " + err.Error())
		return err
	} else if result.DeletedCount == 0 {
		msg := "Failed to delete the Book with id = " + id
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
}
