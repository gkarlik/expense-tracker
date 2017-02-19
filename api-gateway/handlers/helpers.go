package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
)

func ParseRequestData(s quark.Service, r *http.Request, in interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		LogError(s, err, "Cannot read request body")
		return err
	}
	err = json.Unmarshal(b, in)
	if err != nil {
		LogError(s, err, "Cannot parse request body")
		return err
	}
	return nil
}

func JSONResponse(s quark.Service, w http.ResponseWriter, in interface{}) error {
	data, err := json.Marshal(in)
	if err != nil {
		LogError(s, err, "Cannot send JSON reponse")
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return nil
}

func LogError(s quark.Service, e error, message string) {
	s.Log().ErrorWithFields(logger.Fields{"error": e}, message)
}
