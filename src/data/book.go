package data

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Year      int                `json:"year,omitempty" bson:"year,omitempty"`
	ShortDesc string             `json:"shortDesc,omitempty" bson:"shortDesc,omitempty"`
	Genre     string             `json:"genre,omitempty" bson:"genre,omitempty"`
}

// Books = collection of Book
type Books []*Book

//conversion helper methods
func (b *Book) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}

/*"ToJSON serializes the contents of the collection to JSON
NewEncoder provides better performance than json.Unmarshal() as it does not
have to buffer the output into an in memory slice of bytes
this reduces allocations and the overheads of the service"*/
func (b *Books) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

//Database interaction methods
func GetBooks(bc *mongo.Collection) (Books, error) {
	var bookList Books

	//get all books from the database
	cur, err := bc.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("find operation on the db collection: " + err.Error())
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var b *Book
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

/*Accepts id of type primitive.ObjectId and database collection. Returns *Book and error.*/
func GetBookById(id primitive.ObjectID, bc *mongo.Collection) (*Book, error) {
	var b Book
	filter := bson.M{"_id": id}

	//check if the result isn't null (no matching documents were found)
	result := bc.FindOne(context.Background(), filter)
	if result == nil {
		return nil, nil
	}

	err := result.Decode(&b)
	if err != nil {
		log.Println("error while decoding the db results: ", err.Error())
		return nil, err
	}

	return &b, nil
}

/*Takes in the pointers to Book and database collection. Inserts the Book into the database. Returns the id of inserted book and error.*/
func InsertBook(b *Book, bc *mongo.Collection) (string, error) {
	var insertedId string
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := bc.InsertOne(ctx, b)
	if err != nil {
		return "", err
	}
	insertedId = result.InsertedID.(primitive.ObjectID).Hex() //converts primitive.ObjectId to string
	return insertedId, err
}

func UpdateBook(b Book, bc *mongo.Collection) error {
	filter := bson.M{"_id": b.Id} //filter := bson.M{"_id": bson.M{"$eq": objId}}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := bc.UpdateOne(ctx, filter, bson.D{{"$set", b}}) //because of omitempty, only the fields we fill in are updated
	if err != nil {
		log.Println("performing an update operation: " + err.Error())
		return err
	} else if result.ModifiedCount == 0 {
		msg := "Failed to update the Book with id = " + b.Id.Hex()
		log.Println(msg)
		return errors.New(msg)
	}
	return nil
}

/*Takes in objectId Id and database collection. Deletes the books with this id from the database, returns error.*/
func DeleteBook(id primitive.ObjectID, bc *mongo.Collection) error {
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := bc.DeleteOne(ctx, filter)

	if err != nil {
		log.Println("performing a delete operation: " + err.Error())
		return err
	} else if result.DeletedCount == 0 {
		msg := "Failed to delete the Book with id = " + id.Hex()
		log.Println(msg)
		return errors.New(msg)
	}

	return nil
}
