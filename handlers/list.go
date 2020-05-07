package handlers

import (
	"log"
	"net/http"

	contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/rest-test/lib"
)

func List(w http.ResponseWriter, r *http.Request) {
	p := contacts.LoaPersistent()
	list, err := p.List()
	if err != nil {
		log.Fatal(err)
	}
	lib.ReturnJSON(w, list)
}
