package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgryski/go-trigram"
	contacts "github.com/wshaman/contacts-stub"
	"github.com/wshaman/course-rest/lib"
)

type SearchRequest = struct {
	Text string `json:"text,omitempty"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	p := contacts.LoadPersistent()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	schReq := SearchRequest{}
	if err = json.Unmarshal(b, &schReq); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	existing, _ := p.List()
	var rtn []contacts.Contact
	for _, c := range existing {
		var searchGroup []string
		searchGroup = append(searchGroup, c.FirstName, c.LastName, c.Phone)
		idx := trigram.NewIndex(searchGroup)
		found := idx.Query(schReq.Text)
		if len(found) > 0 {
			rtn = append(rtn, c)
		}
	}
	if err != nil {
		lib.ReturnInternalError(w, err)
	}
	lib.ReturnJSON(w, rtn)
}
