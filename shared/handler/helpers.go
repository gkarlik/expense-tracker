package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/middleware/auth/jwt"
	uuid "github.com/satori/go.uuid"
)

const (
	CategoriesPageSize = 50
	ExpensesPageSize   = 50
	UserClaimsKey      = "USER_KEY"
	UserIDKey          = "UserID"
	RequestIDKey       = "Request-ID"
	ErrorKey           = "Error"
	ErrorMetricName    = "errors"
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

	fmt.Println(r.Context())

	reqID, ok := r.Context().Value(RequestIDKey).(string)
	if !ok {
		return ""
	}
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

func setResponse(w http.ResponseWriter, data []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func Response(s quark.Service, w http.ResponseWriter, r *http.Request, in interface{}, code int) {
	data, err := json.Marshal(in)
	if err != nil {
		ErrorResponse(s, w, r, errors.ErrInternal, http.StatusInternalServerError)
		return
	}
	setResponse(w, data, code)
}

func ErrorResponse(s quark.Service, w http.ResponseWriter, r *http.Request, e *errors.Error, code int) {
	quark.ReportError(s, r, ErrorKey, ErrorMetricName, e)

	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	setResponse(w, data, code)
}
