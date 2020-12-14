package book

type Book struct {
	Id        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	ShortDesc string `json:"shortDesc" bson:"shortDesc"` //~50symbols
	Year      int    `json:"year" bson:"year"`
	Genre     string `json:"genre" bson:"genre"`
	//imgUrl stirng//
	//price float //retrieve from amazon api, look it up
	//rating float
}
