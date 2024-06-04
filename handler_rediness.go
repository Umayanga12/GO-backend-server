package main

import (
	"net/http"
)

// HandlerRediness - check the rediness of the server - health check
// weaher the server is ready to serve the request or not
func HandlerRediness(w http.ResponseWriter, r *http.Request) {
	responceWithJson(w, 200, struct{}{})
}
