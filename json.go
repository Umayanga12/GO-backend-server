package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
