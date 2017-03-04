package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/middleware/auth/jwt"
	em "github.com/gkarlik/quark-go/middleware/error"
	"github.com/gkarlik/quark-go/middleware/logging"
	mm "github.com/gkarlik/quark-go/middleware/metrics"
	"github.com/gkarlik/quark-go/middleware/ratelimiter"
	"github.com/gkarlik/quark-go/middleware/security"
	tm "github.com/gkarlik/quark-go/middleware/tracing"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var authenticationMiddlewareHandler = auth.NewAuthenticationMiddleware(
	auth.WithSecret(quark.GetEnvVar("GATEWAY_SECRET")),
	auth.WithContextKey(handler.UserClaimsKey),
	auth.WithAuthenticationFunc(func(credentials auth.Credentials) (auth.Claims, error) {
		return auth.Claims{}, errors.New("Authenticate function is not set")
	}),
)

func CreateLimiterMiddleware(s quark.Service) *negroni.Negroni {
	var limterMiddleware = negroni.HandlerFunc(ratelimiter.NewRateLimiterMiddleware(10 * time.Millisecond).HandleWithNext)
	var logRequestMiddleware = negroni.HandlerFunc(logging.NewRequestLoggingMiddleware(handler.RequestIDKey).HandleWithNext)
	var errorMiddleware = negroni.HandlerFunc(em.NewRequestErrorMiddleware().HandleWithNext)
	var metricsMiddleware = negroni.HandlerFunc(mm.NewRequestMetricsMiddleware(s).HandleWithNext)
	var tracingMiddleware = negroni.HandlerFunc(tm.NewRequestTracingMiddleware(s).HandleWithNext)
	var securityMiddleware = negroni.HandlerFunc(security.NewRequestSecurityMiddleware().HandleWithNext)

	return negroni.New(
		metricsMiddleware,
		errorMiddleware,
		securityMiddleware,
		tracingMiddleware,
		logRequestMiddleware,
		limterMiddleware,
	)
}

func CreateCommonMiddleware(s quark.Service) *negroni.Negroni {
	var authenticationMiddleware = negroni.HandlerFunc(authenticationMiddlewareHandler.AuthenticateWithNext)
	limiterMiddleware := CreateLimiterMiddleware(s)

	return limiterMiddleware.With(
		authenticationMiddleware,
	)
}

func Init(r *mux.Router, s quark.Service) *mux.Router {
	authenticationMiddlewareHandler.Options.Authenticate = func(credentials auth.Credentials) (auth.Claims, error) {
		return v1.AuthenticateUser(s, credentials)
	}

	api := r.PathPrefix("/api/v1").Subrouter()

	limterMiddleware := CreateLimiterMiddleware(s)

	api.Handle("/auth", limterMiddleware.With(
		negroni.Wrap(http.HandlerFunc(authenticationMiddlewareHandler.GenerateToken))),
	).Methods(http.MethodPost)

	return api
}
