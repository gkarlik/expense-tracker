package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitUsersRoutes(api *mux.Router, s quark.Service) {
	commonMiddleware := CreateCommonMiddleware(s)
	limterMiddleware := CreateLimiterMiddleware(s)

	api.Handle("/users", commonMiddleware.With(
		negroni.Wrap(v1.GetUserByLoginHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/users/{id}", commonMiddleware.With(
		negroni.Wrap(v1.GetUserByIDHandler(s))),
	).Methods(http.MethodGet)

	api.Handle("/users", limterMiddleware.With(
		negroni.Wrap(v1.RegisterUserHandler(s))),
	).Methods(http.MethodPost)

	api.Handle("/users", limterMiddleware.With(
		negroni.Wrap(v1.UpdateUserHandler(s))),
	).Methods(http.MethodPut)
}
