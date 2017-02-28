package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitExpensesRoutes(api *mux.Router, s quark.Service) {
	api.Handle("/expenses", commonMiddleware.With(
		negroni.Wrap(v1.GetExpensesHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/expenses", commonMiddleware.With(
		negroni.Wrap(v1.UpdateExpenseHandler(s))),
	).Methods(http.MethodPut, http.MethodPost)

	api.Handle("/expenses/{id}", commonMiddleware.With(
		negroni.Wrap(v1.GetExpenseHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/expenses/{id}", commonMiddleware.With(
		negroni.Wrap(v1.RemoveExpenseHandler(s))),
	).Methods(http.MethodDelete)
}
