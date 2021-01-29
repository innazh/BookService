package handlers

import (
	"BooksWebservice/data"
	"BooksWebservice/middleware"
	"net/http"
)

/*Handler for /signup POST route. Retrieves user from the context of the request, inserts it to the database.
Sends back StatusCreated header along with the user id.*/
func (e *Env) SignUp(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Handle POST on /signup")
	var err error
	user := r.Context().Value(middleware.KeyUser{}).(data.User) //user object is validated via middleware

	var userInsertId string
	// hash user's password
	if err := data.PrepareUserInsert(&user); err != nil {
		e.l.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// add the user into the database
	if userInsertId, err = data.InsertNewUser(&user, e.userCollection); err != nil {
		e.l.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e.l.Println("Successfully registered a new user. Id=", userInsertId)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(userInsertId))
	return
}
