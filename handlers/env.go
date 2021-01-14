package handlers

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

//collection name for User objects in the database
var userCollectionName string = "users"

//collection name for Book objects in the database
var bookCollectionName string = "books"

/*Maybe this should be in utils*/
/*Env struct contains the dependencies shared by
all handlers of the application such as logger and db connection*/
type Env struct {
	l              *log.Logger
	db             *mongo.Client
	userCollection *mongo.Collection
	bookCollection *mongo.Collection
}

func NewEnv(l *log.Logger, db *mongo.Client, dbName string) *Env {
	return &Env{l, db, db.Database(dbName).Collection(userCollectionName), db.Database(dbName).Collection(bookCollectionName)}
}

// //this seems more proper but a little extra
// func (e *Env) GetUserCollection() *mongo.Collection {
// 	return db.Database(dbName).Collection(userCollectionName)
// }

// func GetBookCollection() {
// 	return db.Database(dbName).Collection(bookCollectionName)
// }
