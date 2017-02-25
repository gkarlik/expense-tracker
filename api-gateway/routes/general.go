package routes

import (
	"time"

	"errors"
	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/ratelimiter"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

var limterMiddlewareHandler = negroni.HandlerFunc(
	ratelimiter.NewHTTPRateLimiter(100 * time.Millisecond).HandleWithNext,
)

var limterMiddleware = negroni.New(
	limterMiddlewareHandler,
)

var authenticationMiddlewareHandler = auth.NewAuthenticationMiddleware(
	auth.WithSecret(quark.GetEnvVar("GATEWAY_SECRET")),
	auth.WithContextKey("USER_KEY"),
	auth.WithAuthenticationFunc(func(credentials auth.Credentials) (auth.Claims, error) {
		return auth.Claims{}, errors.New("Authenticate function is not set")
	}),
)

var authenticationMiddleware = negroni.HandlerFunc(authenticationMiddlewareHandler.AuthenticateWithNext)

var commonMiddleware = negroni.New(
	limterMiddlewareHandler,
	authenticationMiddleware,
)

func Init(r *mux.Router, s quark.Service) *mux.Router {
	authenticationMiddlewareHandler.Options.Authenticate = func(credentials auth.Credentials) (auth.Claims, error) {
		return v1.AuthenticateUser(s, credentials)
	}

	api := r.PathPrefix("/api/v1").Subrouter()

	api.Handle("/auth", limterMiddleware.With(
		negroni.Wrap(http.HandlerFunc(authenticationMiddlewareHandler.GenerateToken))),
	).Methods(http.MethodPost)

	return api
}
