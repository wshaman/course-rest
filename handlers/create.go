package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/rest-test/lib"
)

func Create(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoaPersistent()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ReturnInternalError(w)
		return
	}
	contact := stub_contacts.Contact{}
	if err := json.Unmarshal(b, &contact); err != nil {
		lib.ReturnInternalError(w)
		return
	}
	contact.ID = 0
	uid, err := p.Save(contact)
	if err != nil {
		lib.ReturnInternalError(w)
		return
	}
	newContact, err := p.Load(uid)
	if err != nil {
		lib.ReturnInternalError(w)
		return
	}
	lib.ReturnJSON(w, newContact)
}
