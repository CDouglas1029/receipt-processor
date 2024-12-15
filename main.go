package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // Correct import path
)

func main() {
	router := mux.NewRouter()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}