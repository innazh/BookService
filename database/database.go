package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbConn *mongo.Client //*sql.DB //*mongo.Client

func SetupDB() {
	var err error
	DbConn, err = mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DbConn.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//defer DbConn.Disconnect(ctx) //probably comment out and make a function for closing the connwection which we can call fom the main func
}
