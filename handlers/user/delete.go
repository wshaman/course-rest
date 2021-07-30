package user

import (
	"net/http"

	stub_contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/course-rest/lib"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	p := stub_contacts.LoadPersistent()
	id, err := lib.IDFromVars(r)
	if err != nil {
		lib.ReturnClientError(w, err.Error())
		return
	}
	if _, err = p.Load(id); err != nil {
		lib.ReturnClientError(w, "user not found")
		return
	}
	if err := p.Delete(id); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
}
