package main

import (
	"github.com/rodkevich/course-rest/handlers/user"
	stub_contacts "github.com/wshaman/contacts-stub"
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
	router.HandleFunc("/api/v0/user/search", user.Search).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/user/search-cheat", user.CheatSearch).Methods(http.MethodPost)
	log.Println("Populating data")
	p := stub_contacts.LoadPersistent()
	if err := stub_contacts.Populate(p); err != nil {
		log.Fatal(err)
	}
	log.Println("Starting API server on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
