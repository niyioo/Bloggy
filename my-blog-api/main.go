package main

import (
	"my-blog-api/database"
	"my-blog-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the MongoDB database
	client, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Close(client) // Close the client when the main function exits

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Register routes for user and blog post actions
	routes.RegisterUserRoutes(router)

	// Specify the address and port for your API server
	addr := "localhost:8080"

	// Create a server and set some basic configurations
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start the server with more error handling
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
