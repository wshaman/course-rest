package lib

import (
	"log"
	"net/http"
)

func ReturnInternalError(w http.ResponseWriter, errHappened error) {
	log.Println(errHappened.Error())
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`Sorry pal, not this time`)); err != nil {
		log.Fatal(err)
	}
}

func ReturnClientError(w http.ResponseWriter, text string) {
	log.Println(text)
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}
