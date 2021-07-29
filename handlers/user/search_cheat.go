package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rodkevich/course-rest/lib"
	contacts "github.com/wshaman/contacts-stub"
)

func CheatSearch(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	p := contacts.LoadPersistent()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	if err = json.Unmarshal(b, &request); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	rtn, err := p.FindByPhone(request["text"].(string))
	if err != nil {
		lib.ReturnInternalError(w, err)
	}
	lib.ReturnJSON(w, rtn)
}
