package cors

//In Go, middleware is referred to as handlers
import (
	"BooksWebservice/services"
	"fmt"
	"net/http"
	"time"
)

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

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
        c, err := r.Cookie("token")
		
		if err != nil {
                if err == http.ErrNoCookie {
                        // If the cookie is not set, return an unauthorized status
                        w.WriteHeader(http.StatusUnauthorized)
                        return
                }
                // For any other type of error, return a bad request status
                w.WriteHeader(http.StatusBadRequest)
                return
        }

        // Get the JWT string from the cookie
		jwtStr := c.Value
		if err = services.VerifyToken(jwtStr); err!=nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return	
		}
		if !jwtStr.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		














		authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error")
                    }
                    return []byte("secret"), nil
                })
                if error != nil {
                    json.NewEncoder(w).Encode(Exception{Message: error.Error()})
                    return
                }
                if token.Valid {
                    context.Set(req, "decoded", token.Claims)
                    next(, wreq)
                } else {
                    json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
                }
            }
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
        }
    })
}
