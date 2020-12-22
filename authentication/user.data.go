package authentication

import (
	"BooksWebservice/database"
	"BooksWebservice/services"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrepareUserInsert(user *User) error {
	var err error
	if user.Password, err = services.HashAndSalt([]byte(user.Password)); err != nil {
		return err
	}
	return nil
}

func InserNewUser(user User) (string, error) {
	dbColl := database.GetMongoDbCollection("Books", "users")

	result, err := dbColl.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetUserById(id string) (*User, error) {
	// convert given ID to ObjectId
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("converting Id: " + err.Error())
		return nil, err
	}

	var user User

	dbColl := database.GetMongoDbCollection("Books", "users")
	result := dbColl.FindOne(context.Background(), bson.M{"_id": objId})
	if result == nil {
		return nil, nil
	}

	if err = result.Decode(&user); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User

	dbColl := database.GetMongoDbCollection("Books", "users")
	result := dbColl.FindOne(context.Background(), bson.M{"username": username})
	if result == nil {
		return nil, nil
	}

	if err := result.Decode(&user); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

//maybe create a function here to prepare User object to be written into the db
//it'd hash the password and replace the password in the object with hashed password
