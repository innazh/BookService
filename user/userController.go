package user

import (
	"BooksWebservice/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetupRoutes(apiBasePath string) {
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "register"), http.HandlerFunc(RegistrationHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "signin"), http.HandlerFunc(SignInHandler))
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
		if user.Id == primitive.NilObjectID && user.Username != "" && user.Password != "" {
			// hash user's password
			if err = PrepareUserInsert(&user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// add the user into the database
			if userInsertId, err = InserNewUser(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(userInsertId))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return

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

		///get user from the database by username
		dbUser, err := GetUserByUsername(user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if dbUser == nil && err == nil { //user with such ID doesn't exist!
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User such username doesn't exist!"))
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
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//later: ensure usernames are unique.
//ensure: usernames aren't empty
//ensure: passwords aren't empty or are between 6 and 24 symbols
