package handlers

import (
	"BooksWebservice/data"
	"BooksWebservice/middleware"
	"BooksWebservice/utils"
	"net/http"
)

/*Handler for /signin POST route. Retrieves user from the context of the request, validates it against the data in the database.
Sends back a JWT token with 5 minute expiration time*/
func (e *Env) SignIn(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle user Sign-in")
	user := r.Context().Value(middleware.KeyUser{}).(data.User) //we validate user in middleware

	///get user from the database by username
	dbUser, err := data.GetUserByUsername(user.Username, e.userCollection)
	if err != nil {
		e.l.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if dbUser == nil && err == nil { //user with such ID doesn't exist!
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User with such username doesn't exist!"))
		return
	}

	match := utils.ComparePasswords([]byte(dbUser.Password), []byte(user.Password))
	//if the result successful - grant a JWT token to the user, if no - provide an 'unauthorarized access error'
	if match {

		jwtToken, expTime, err := utils.CreateToken(utils.GetKey(), user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//w.WriteHeader(http.StatusOK)
		//w.Write([]byte("User " + user.Username + " is successfully authenticated!"))
		http.SetCookie(w, &http.Cookie{Name: "token", Value: jwtToken, Expires: expTime, HttpOnly: true, SameSite: http.SameSiteStrictMode})

		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong password"))
		return
	}
}
