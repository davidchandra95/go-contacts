package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go-contacts/app"
	"go-contacts/controllers"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts/{id}", controllers.GetContact).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/me/contacts", controllers.GetContacts).Methods("GET") // user/2/contacts

	router.Use(app.JwtAuthentication) // Attach JWT middleware

	port := os.Getenv("PORT") // Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "3000" //localhost
	}

	fmt.Println(port)

	// Important lines for frond-end utilise
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET", "PUT", "DELETE"})

	err := http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)) // Launch the app
	if err != nil {
		fmt.Print(err)
	}
}