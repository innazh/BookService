package main

//curl --local-port 8080 POST -H "Origin: http://localhost:8080" -H "Content-Type: application/json" --verbose -d '{"username": "admin", "password": "admin"}' http://localhost:5000/signin

import (
	"BooksWebservice/data"
	"BooksWebservice/handlers"
	"BooksWebservice/middleware"
	"BooksWebservice/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

/*Something else to add on: maybe make books exportable as a file on /books/download path?*/
func main() {
	config := utils.GetConfiguration()

	l := log.New(os.Stdout, "book-api: ", log.LstdFlags|log.Lshortfile)
	db := data.GetNewClient(config.ConnectionString)
	env := handlers.NewEnv(l, db, config.DbName)

	//Create router and setup routes
	sm := mux.NewRouter().StrictSlash(true)

	//GET
	setupGETRoutes(sm, env)
	//POST
	setupPOSTRoutes(sm, env)
	//PUT
	setupPUTRoutes(sm, env)
	//DELETE
	setupDELETERoutes(sm, env)

	//CORS - allow all origins
	ch := gohandlers.CORS(gohandlers.AllowCredentials(), gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}), gohandlers.AllowedOrigins([]string{"http://127.0.0.1:3000"}), gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}))

	//create a custom server to change the timeouts, port & assign the configured router sm to it
	s := &http.Server{
		Addr:         config.Port,
		Handler:      ch(sm),
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	println("Server is running on localhost" + config.Port)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	serverShutdown(l, s, db, config.DbName)
}

//Sets up all GET route handlers and middleware
func setupGETRoutes(sm *mux.Router, env *handlers.Env) {
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.Use(middleware.ValidateJWTToken) //validate JWT token
	getRouter.Handle("/", http.RedirectHandler("/books", 301))
	getRouter.Handle("/books", http.HandlerFunc(env.GetBooks))
	getRouter.Handle("/books/{id}", http.HandlerFunc(env.GetBook))
	getRouter.Handle("/checksignin", http.HandlerFunc(env.CheckSignIn))
}

//Sets up all POST route handlers and middleware
func setupPOSTRoutes(sm *mux.Router, env *handlers.Env) {
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.PathPrefix("/books").Handler(middleware.ValidateJWTToken(middleware.NewBookValidation(env.AddBook)))
	postRouter.PathPrefix("/signup").Handler(middleware.ValidateNewUser(env.SignUp))
	postRouter.PathPrefix("/signin").Handler(middleware.ValidateUser(env.SignIn))
}

//Sets up all PUT route handlers and middleware
func setupPUTRoutes(sm *mux.Router, env *handlers.Env) {
	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.Use(middleware.ValidateJWTToken, middleware.UpdateBookValidation)
	putRouter.Handle("/books/{id}", http.HandlerFunc(env.UpdateBook))
}

//Sets up all DELETE route handlers and middleware //TODO: add some kind of time/tracking middleware, low priority
func setupDELETERoutes(sm *mux.Router, env *handlers.Env) {
	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.Use(middleware.ValidateJWTToken)
	deleteRouter.Handle("/books/{id}", http.HandlerFunc(env.DeleteBook))
}

//close everything donw upon receiving a shut down command
func serverShutdown(l *log.Logger, s *http.Server, db *mongo.Client, dbName string) {
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel

	l.Println("Shutting the server down->", sig)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
	db.Database(dbName) //what? TODO!
}
