package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/expense-tracker/shared/errors"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	uuid "github.com/satori/go.uuid"
)

const (
	CategoriesPageSize = 50
	ExpensesPageSize   = 50
	UserClaimsKey      = "USER_KEY"
	UserIDKey          = "UserID"
	RequestIDKey       = "Request-ID"
)

func GetUserID(r *http.Request) string {
	if r == nil {
		return ""
	}

	claims := r.Context().Value(UserClaimsKey).(auth.Claims)
	userID := claims.Properties[UserIDKey].(string)

	return userID
}

func GetRequestID(r *http.Request) string {
	if r == nil {
		return uuid.NewV4().String()
	}

	reqID := r.Context().Value(RequestIDKey).(string)

	return reqID
}

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

func Response(w http.ResponseWriter, in interface{}, code int) {
	data, err := json.Marshal(in)
	if err != nil {
		ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
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
