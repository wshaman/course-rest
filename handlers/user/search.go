package user

import (
	"net/http"

	contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/course-rest/lib"
)

func Search(w http.ResponseWriter, r *http.Request) {
	// url transfer as a parameter
	query := r.URL.Query()
	number := query.Get("number")
	p := contacts.LoadPersistent()
	// find by phone
	res, err := p.FindByPhone(number)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	lib.ReturnJSON(w, res)
}