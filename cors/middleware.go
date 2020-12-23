package cors

//In Go, middleware is referred to as handlers
import (
	"BooksWebservice/services"
	"BooksWebservice/settings"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*A regular middleware func that sets headers and fixes the amount of time it took for request to complete*/
func MiddlewareFunc(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//CORS headers:
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//
		fmt.Println("before handler; middlerware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middlerware finished; %s\n", time.Since(start))
	})
}

/*Function checks for a jwt token in a Cookie of a request, varifies it and ServesHTTP
If the token isn't valid, sends back StatusUnauthorized header*/
func ValidateMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtStr := c.Value
		token, err := services.VerifyToken(settings.GetKey(), jwtStr)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
