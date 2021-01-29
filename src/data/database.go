package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Connects to the remote db, returns a db client instance*/
//mongodb+srv://<username>:<password>@cluster0.enzda.mongodb.net/<dbname>?retryWrites=true
func GetNewClient(connStr string) *mongo.Client {
	var err error
	DbConn, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DbConn.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return DbConn
}
