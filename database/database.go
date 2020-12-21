package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbConn *mongo.Client //*sql.DB //*mongo.Client

func SetupDB(conn_str string) {
	var err error
	DbConn, err = mongo.NewClient(options.Client().ApplyURI(conn_str))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DbConn.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMongoDbCollection(dbName, collName string) *mongo.Collection {
	return DbConn.Database(dbName).Collection(collName)
}

func CloseDbConn() {
	DbConn.Disconnect(context.TODO())
}
