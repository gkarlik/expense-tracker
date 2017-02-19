package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
)

func ParseRequestData(r *http.Request, in interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, in)
	if err != nil {
		return err
	}
	return nil
}

func LogError(s quark.Service, e error, message string) {
	s.Log().ErrorWithFields(logger.Fields{"error": e}, message)
}
