package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(apiBasePath string) {
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "/register"), http.HandlerFunc(RegistrationHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, "/signin"), http.HandlerFunc(SignInHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, "/refresh"), http.HandlerFunc())
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var userInsertId string
		//if the sent object is valid - send it off into the database to encode and store
		if &user != nil && user.id == primitive.NilObjectID && user.username != "" && user.password != "" {
			//TODO: hash user's password
			if err = PrepareUserInsert(&user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			//TODO: add the user into the database
			if userInsertId, err = InserNewUser(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(userInsertId))
			return
			//think about the fields you might wanna have here.........
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//TODO: get user from the database by ID
		if dbUser, err := GetUserById(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		//TODO: get the retrieved user's hashed password and convert it into a []byte
		//TODO: call comparePasswords() from passwordManager and deal with the result:
		//if the result successful - grant a JWT token to the user, if no - provide an 'unauthorarized access error'
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
