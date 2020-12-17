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

//https://medium.com/better-programming/building-a-restful-api-with-go-and-mongodb-93e59cbbee88
// func GetCollection(string collName){
// 	return DbConn.Database("Books")
// }
//not bad error handling: https://medium.com/@faygun89/create-rest-api-with-golang-and-mongodb-d38d2e1d9714

func CloseDbConn() {
	DbConn.Disconnect(context.TODO())
}
