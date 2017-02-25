package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitCategoriesRoutes(parent *mux.Router, s quark.Service) {
	categories := parent.PathPrefix("/categories").Subrouter()

	categories.Handle("/", commonMiddleware.With(
		negroni.Wrap(v1.UpdateCategoryHandler(s))),
	).Methods(http.MethodPut, http.MethodPost)
}
