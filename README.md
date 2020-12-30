# BookService - is an API
BookService is developed in Go using net/http package and connected to MongoDB. Athorization to use the service is granted through JWT tokens with a lifespane of 5 minutes. JWT tokens are granted to registered users upon the sign in. 
The BookService supports GET, PUT, DELETE and POST requests for registered users.

## API usage

### User
Operations related to user.
#### /user/register - POST
User must register to be able to use the API. 
Send a JSON object containing "username" and "password" fields in a POST request to be registered. 
Example:
```JSON
{
    "username":"username",
    "password":"password"
}
```
#### /user/signin - POST
After registering, authenticate by sending a POST request containing your username and password.
Upon successful authentication, you'll receive a JWT token with a set expiration time

### Books
In order to have the access to the core routes of the API, user must be authenticated and hold a valid JWT token in a Cookie.
#### Book model
```JSON
{
    "id": "id",
    "title": "title",
    "author": "author",
    "year": 1990,
    "shortDesc": "description",
    "genre": "comma, separated, list, of, genres"
}
```
#### /api/books - GET
Gets all available books from the database in JSON format.

#### /api/books - POST
Adds a book passed via this request to the database. Book's id must be null or empty string.
Returns an Id of the inserted book if everything is well.

#### /api/books/{id} - GET
Gets a book in with the requested Id in JSON format.

#### /api/books/{id} - UPDATE
To update a book, send along a JSON object with only those fields that you wish to update. For example, if I want to update only description and year, I'll send:
```JSON
{
    "year": 1990,
    "shortDesc": "description"
}
```
You can't update the Id of the book.

#### /api/books/{id} - DELETE
The book with this id gets deleted from the database.
