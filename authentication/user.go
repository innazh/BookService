package authentication

import "go.mongodb.org/mongo-driver/bson/primitive"

//model to be extracted from the request body
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}
