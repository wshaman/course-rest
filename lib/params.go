package lib

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func IDFromVars(r *http.Request) (uint, error) {
	idString := mux.Vars(r)["id"]
	i, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func FieldsFromStruct(i interface{}) (res []string) {
	v := reflect.ValueOf(i)
	for j := 0; j < v.NumField(); j++ {
		res = append(res, v.Field(j).String())
	}
	return
}
