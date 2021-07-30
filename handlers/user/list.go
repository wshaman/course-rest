package user

import (
	"net/http"

	contacts "github.com/wshaman/contacts-stub"

	"github.com/wshaman/course-rest/lib"
)

func List(w http.ResponseWriter, r *http.Request) {
	p := contacts.LoadPersistent()
	list, err := p.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, list)
}
