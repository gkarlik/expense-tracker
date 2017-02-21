package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/expense-tracker/shared/errors"
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

func Response(w http.ResponseWriter, in interface{}) {
	data, err := json.Marshal(in)
	if err != nil {
		ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, e *errors.Error, code int) {
	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(data)
}
