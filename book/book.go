package book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author"`
	Year      int                `json:"year" bson:"year"`
	ShortDesc string             `json:"shortDesc" bson:"shortDesc"` //~50symbols
	Genre     string             `json:"genre" bson:"genre"`
	//imgUrl stirng//
	//price float //retrieve from amazon api, look it up
	//rating float
}
