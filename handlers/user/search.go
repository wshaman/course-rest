package user

import (
	"encoding/json"
	"fmt"
	"github.com/rodkevich/course-rest/lib"
	contacts "github.com/wshaman/contacts-stub"
	"io/ioutil"
	"net/http"
	"sort"
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

	sr := SearchRequest{}
	if err = json.Unmarshal(b, &sr); err != nil {
		lib.ReturnInternalError(w, err)
		return
	}

	existing, _ := p.List()
	var rtn []contacts.Contact
	for _, c := range existing {
		var searchGroup []string
		searchGroup = append(searchGroup, c.FirstName, c.LastName, c.Phone)
		if pos := sort.SearchStrings(searchGroup, sr.Text); pos < len(searchGroup) {
			rtn = append(rtn, c)
		}
	}
	fmt.Printf("%q\n", rtn)
	if err != nil {
		lib.ReturnInternalError(w, err)
	}
	lib.ReturnJSON(w, rtn)
}
