package user

import (
	"fmt"
	"log"
	"net/http"

	contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/course-rest/lib"
)

func Find(w http.ResponseWriter, r *http.Request) {
	a := contacts.LoadPersistent()
	list, err := a.List()
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	id, err := lib.IDFromVars(r)
	if err != nil {
		lib.ReturnClientError(w, err.Error())
		return
	}

	if _, err = a.Load(id); err != nil {
		message := fmt.Sprintf("Contact %d not found.", id)
		if _, err := w.Write([]byte(message)); err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	lib.ReturnJSON(w, list[id])
}