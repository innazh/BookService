package data

import (
	"BooksWebservice/utils"
	"context"
	"encoding/json"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//model to be extracted from the request body
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}

//conversion helper methods
func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

//Database interaction methods
/*Prepares user object for the insertion into the database - hashes and salts their password*/
func PrepareUserInsert(user *User) error {
	var err error
	if user.Password, err = utils.HashAndSalt([]byte(user.Password)); err != nil {
		return err
	}
	return nil
}

/*Takes in user and database collection. Inserts user in the database. Returns the Id of the inserted user and an error object*/
func InsertNewUser(user *User, uc *mongo.Collection) (string, error) {
	result, err := uc.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

/*WORK IN PROGRESS, unused by the code for now
Accepts object id and database client. Returns the user with that id and error object*/
func GetUserById(id primitive.ObjectID, uc *mongo.Collection) (*User, error) {
	// convert given ID to ObjectId
	// objId, err := primitive.ObjectIDFromHex(id)//converts string to object Id
	// if err != nil {
	// log.Println("converting Id: " + err.Error())
	// return nil, err
	// }

	var user User

	// dbCol := db.Database("Books").Collection("users")
	// result := dbCol.FindOne(context.Background(), bson.M{"_id": objId})
	// if result == nil {
	// return nil, nil
	// }

	// if err = result.Decode(&user); err != nil {
	// log.Println(err.Error())
	// return nil, err
	// }

	return &user, nil
}

/*Takes in username and database collection. Returns a user object and an error.*/
func GetUserByUsername(username string, uc *mongo.Collection) (*User, error) {
	var user User

	result := uc.FindOne(context.Background(), bson.M{"username": username})
	if result == nil {
		return nil, nil
	}

	if err := result.Decode(&user); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}
