package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wshaman/rest-test/handlers"

)


func main () {
	router := mux.NewRouter()
	router.HandleFunc("/api/v0/user", handlers.List)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
