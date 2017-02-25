package routes

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func InitUsersRoutes(parent *mux.Router, s quark.Service) {
	users := parent.PathPrefix("/users").Subrouter()

	users.Handle("/", commonMiddleware.With(
		negroni.Wrap(v1.GetUserByLoginHandler(s))),
	).Methods(http.MethodGet)

	users.Handle("/{id}/", commonMiddleware.With(
		negroni.Wrap(v1.GetUserByIDHandler(s))),
	).Methods(http.MethodGet)

	users.Handle("/", limterMiddleware.With(
		negroni.Wrap(v1.RegisterUserHandler(s))),
	).Methods(http.MethodPost)
}
