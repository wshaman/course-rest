package user

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/course-rest/lib"
)

func Search(w http.ResponseWriter, r *http.Request) {
	p := contacts.LoadPersistent()
	list, err := p.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	phoneString := mux.Vars(r)["phoneNumber"]
	f := true
	for _, v := range list {
		if strings.Contains(v.Phone, phoneString) {
			user, err := p.Load(v.ID)
			if err != nil {
				lib.ReturnInternalError(w, err)
				return
			}
			lib.ReturnJSON(w, user)
			f = fasle
		}
	}
	if f {
		lib.ReturnClientError(w, "No one users contains "+phoneString+" in phone")
	}
}
