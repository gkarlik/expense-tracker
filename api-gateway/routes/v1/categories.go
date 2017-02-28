package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitCategoriesRoutes(api *mux.Router, s quark.Service) {
	api.Handle("/categories", commonMiddleware.With(
		negroni.Wrap(v1.GetCategoriesHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/categories", commonMiddleware.With(
		negroni.Wrap(v1.UpdateCategoryHandler(s))),
	).Methods(http.MethodPut, http.MethodPost)

	api.Handle("/categories/{id}", commonMiddleware.With(
		negroni.Wrap(v1.GetCategoryHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/categories/{id}", commonMiddleware.With(
		negroni.Wrap(v1.RemoveCategoryHandler(s))),
	).Methods(http.MethodDelete)
}