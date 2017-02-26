package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitExpensesRoutes(parent *mux.Router, s quark.Service) {
	expenses := parent.PathPrefix("/expenses").Subrouter()

	expenses.Handle("/", commonMiddleware.With(
		negroni.Wrap(v1.UpdateExpenseHandler(s))),
	).Methods(http.MethodPut, http.MethodPost)
}