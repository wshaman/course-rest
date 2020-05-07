package handlers

import (
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/rest-test/lib"
)

func View(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoaPersistent()
	id, err := lib.IDFromVars(r)
	if err != nil {
		lib.ReturnClientError(w, err.Error())
		return
	}
	user, err := p.Load(id)
	if err != nil {
		lib.ReturnClientError(w, "user not found")
		return
	}
	lib.ReturnJSON(w, user)
}
