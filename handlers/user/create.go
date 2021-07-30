package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/course-rest/lib"
)

func Create(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoadPersistent()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	contact := stub_contacts.Contact{}
	if err := json.Unmarshal(b, &contact); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	contact.ID = 0
	uid, err := p.Save(contact)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	newContact, err := p.Load(uid)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, newContact)
}
