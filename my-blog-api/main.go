package main

import (
	"log"
	"my-blog-api/database"
	"my-blog-api/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	client, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close(client)

	c := cors.Default()

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	routes.RegisterBlogPostRoutes(router)

	handler := c.Handler(router)

	addr := "localhost:8080"

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Printf("Server is starting at %s...", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
