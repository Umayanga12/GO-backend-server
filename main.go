package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	//lead the .env file - database configuration
	godotenv.Load(".env")

	//get the port from the environment variable
	dbPort := os.Getenv("PORT")

	//check weather the port is set or not
	if dbPort == "" {
		log.Fatal("$PORT must be set")
	}
	//fmt.Println("PORT:", dbPort)

	//setup the router
	Router := chi.NewRouter()

	//setup the cors - allow to make request from any origin/ browsers
	Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{""},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//setup the routes
	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", HandlerRediness)

	Router.Mount("/v1", v1Router)

	//connetct the router into the server
	server := &http.Server{
		Addr:    ":" + dbPort,
		Handler: v1Router,
	}

	fmt.Println("Server is running on port: ", dbPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
