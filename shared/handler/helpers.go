package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/expense-tracker/shared/errors"
)

const (
	CategoriesPageSize = 50
	ExpensesPageSize   = 50
)

func ParseRequestData(r *http.Request, in interface{}) ([]byte, error) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return b, err
	}

	err = json.Unmarshal(b, in)
	if err != nil {
		return b, err
	}
	return b, nil
}

func Response(w http.ResponseWriter, r *http.Request, in interface{}) {
	data, err := json.Marshal(in)
	if err != nil {
		ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
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
