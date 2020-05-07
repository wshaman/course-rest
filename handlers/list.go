package handlers

import (
	"net/http"

	//contacts "github.com/wshaman/contacts-stub"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from List handler"))
	//contacts.LoaPersistent()
}
