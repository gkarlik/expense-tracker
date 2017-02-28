package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/ratelimiter"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/urfave/negroni"
	"golang.org/x/net/context"
)

var limterMiddlewareHandler = negroni.HandlerFunc(
	ratelimiter.NewHTTPRateLimiter(100 * time.Millisecond).HandleWithNext,
)

var limterMiddleware = negroni.New(
	requestTraceMiddleware,
	limterMiddlewareHandler,
)

var authenticationMiddlewareHandler = auth.NewAuthenticationMiddleware(
	auth.WithSecret(quark.GetEnvVar("GATEWAY_SECRET")),
	auth.WithContextKey(handler.UserClaimsKey),
	auth.WithAuthenticationFunc(func(credentials auth.Credentials) (auth.Claims, error) {
		return auth.Claims{}, errors.New("Authenticate function is not set")
	}),
)

var authenticationMiddleware = negroni.HandlerFunc(authenticationMiddlewareHandler.AuthenticateWithNext)

var requestTraceMiddleware = negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	reqID := uuid.NewV4()
	ctx := context.WithValue(r.Context(), handler.RequestIDKey, reqID.String())
	r = r.WithContext(ctx)

	logger.Log().DebugWithFields(logger.Fields{"requestID": reqID, "request": r}, "Received request")

	next(w, r)
})

var commonMiddleware = negroni.New(
	requestTraceMiddleware,
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
