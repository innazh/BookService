package authentication

import "go.mongodb.org/mongo-driver/bson/primitive"

//model to be extracted from the request body
type User struct {
	id       primitive.ObjectID
	username string
	password string
}
