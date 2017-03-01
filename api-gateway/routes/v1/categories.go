package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitCategoriesRoutes(api *mux.Router, s quark.Service) {
	commonMiddleware := CreateCommonMiddleware(s)

	api.Handle("/categories", commonMiddleware.With(
		negroni.Wrap(v1.GetCategoriesHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/categories", commonMiddleware.With(
		negroni.Wrap(v1.CreateCategoryHandler(s))),
	).Methods(http.MethodPost)

	api.Handle("/categories", commonMiddleware.With(
		negroni.Wrap(v1.UpdateCategoryHandler(s))),
	).Methods(http.MethodPut)

	api.Handle("/categories/{id}", commonMiddleware.With(
		negroni.Wrap(v1.GetCategoryHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/categories/{id}", commonMiddleware.With(
		negroni.Wrap(v1.RemoveCategoryHandler(s))),
	).Methods(http.MethodDelete)
}
