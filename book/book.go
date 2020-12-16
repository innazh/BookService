package book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Year      int                `json:"year,omitempty" bson:"year,omitempty"`
	ShortDesc string             `json:"shortDesc,omitempty" bson:"shortDesc,omitempty"` //~50symbols
	Genre     string             `json:"genre,omitempty" bson:"genre,omitempty"`
	//imgUrl stirng//
	//price float //retrieve from amazon api, look it up
	//rating float
}
