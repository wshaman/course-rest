package main

import (
	"github.com/wshaman/course-rest/handlers/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/user", user.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user", user.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/user/{id}", user.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v0/user/{id}", user.View).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user/{id}", user.Update).Methods(http.MethodPut)
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
