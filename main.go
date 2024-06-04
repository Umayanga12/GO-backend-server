package main

import (
	"encoding/json"
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
		Handler: Router,
	}

	//server is running on the port
	fmt.Println("Server is running on port: ", dbPort)
	//display the error if any
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

// HandlerRediness - check the rediness of the server - health check
// weaher the server is ready to serve the request or not
func HandlerRediness(w http.ResponseWriter, r *http.Request) {
	responceWithJson(w, 200, struct{}{})
}

// responceWithJson - send the json response
func responceWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err, payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//adding application header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// responceWithError - send the error response
func responceWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponce struct {
		ErrorMsg string `json:"message"`
	}

	responceWithJson(w, code, errResponce{
		ErrorMsg: msg,
	})
}
