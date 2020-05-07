package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wshaman/rest-test/handlers"
)


func main () {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/user", handlers.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user/{id}", handlers.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v0/user/{id}", handlers.View).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user/{id}", handlers.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/v0/user/{id}", handlers.Create).Methods(http.MethodPost)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
