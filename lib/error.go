package lib

import (
	"log"
	"net/http"
)

func ReturnInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`Sorry pal, not this time`)); err != nil {
		log.Fatal(err)
	}
}
