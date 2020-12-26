# BookService - a RESTful API service
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
After registering, authenticate by sending a POST request containing your username and password again.
Upon successful authentication, you'll receive 
