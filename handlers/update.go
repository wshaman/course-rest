package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/rest-test/lib"
)

func Update(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoaPersistent()
	id, err := lib.IDFromVars(r)
	if err != nil {
		lib.ReturnClientError(w, err.Error())
		return
	}
	//contact, err := p.Load(id)
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
	contact.ID = id
	if _, err := p.Save(contact); err != nil {
		lib.ReturnInternalError(w)
		return
	}
	updatedContact, err := p.Load(id)
	if err != nil {
		lib.ReturnInternalError(w)
		return
	}
	lib.ReturnJSON(w, updatedContact)
}
