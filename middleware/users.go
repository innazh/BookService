package middleware

import (
	"BooksWebservice/data"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyUser struct{}

/*Validates a User object upon registration. Checks that the object sent is a valid User object that contains username and password. Password has to pass the requirements.
Validation fails->returns a meaningful error back to the user with BadRequest header. Attaches the user object to the context of the request and calls the next handler
if everything is well.*/
func ValidateNewUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := user.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Invalid user information.", http.StatusBadRequest)
			return
		} else if user.Id == primitive.NilObjectID {
			http.Error(w, "User ids cannot be assigned", http.StatusBadRequest)
		} else if user.Username == "" {
			http.Error(w, "Username is missing", http.StatusBadRequest)
		} else if user.Password == "" {
			http.Error(w, "Password is missing", http.StatusBadRequest)
		} else if len(user.Password) < 6 {
			http.Error(w, "Password must contain 6+ characters", http.StatusBadRequest)
		}

		//pass user with context of the request
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

/*Validates that a JSON object that was sent in the request body is a valid User object.
If Yes->attaches the object to the context of the request
No->Returns BadRequest header back to the user*/
func ValidateUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		err := user.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Invalid user information.", http.StatusBadRequest)
			return
		} else if user == (data.User{}) {
			http.Error(w, "Please provide username and password", http.StatusBadRequest)
		}

		//pass user with context of the request
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
