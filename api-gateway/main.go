package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers/v1"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/metrics/noop"
	"github.com/gkarlik/quark-go/ratelimiter"
	"github.com/gkarlik/quark-go/service/discovery/plain"
	nt "github.com/gkarlik/quark-go/service/trace/noop"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	discovery     *plain.ServiceDiscovery
	discoveryAddr string
)

type Gateway struct {
	*quark.ServiceBase
}

func CreateGateway() *Gateway {
	name := quark.GetEnvVar("GATEWAY_NAME")
	version := quark.GetEnvVar("GATEWAY_VERSION")
	gp := quark.GetEnvVar("GATEWAY_PORT")
	discoveryAddr = quark.GetEnvVar("DISCOVERY")

	port, err := strconv.Atoi(gp)
	if err != nil {
		panic("Incorrect port value!")
	}

	addr, err := quark.GetHostAddress(port)
	if err != nil {
		panic("Cannot resolve host address!")
	}

	discovery = plain.NewServiceDiscovery("http://" + discoveryAddr)

	return &Gateway{
		ServiceBase: quark.NewService(
			quark.Name(name),
			quark.Version(version),
			quark.Address(addr),
			quark.Discovery(discovery),
			quark.Metrics(noop.NewMetricsReporter()),
			quark.Tracer(nt.NewTracer())),
	}
}

var srv = CreateGateway()

func main() {
	defer srv.Dispose()

	secret := quark.GetEnvVar("GATEWAY_SECRET")
	am := auth.NewAuthenticationMiddleware(
		auth.WithSecret(secret),
		auth.WithContextKey("USER_KEY"),
		auth.WithAuthenticationFunc(func(credentials auth.Credentials) (auth.Claims, error) {
			return v1.AuthenticateUser(srv, credentials)
		}))

	rl := ratelimiter.NewHTTPRateLimiter(100 * time.Millisecond)

	r := mux.NewRouter()
	// user service - v1
	r.Handle("/api/v1/auth", rl.Handle(http.HandlerFunc(am.GenerateToken))).Methods(http.MethodPost)
	r.Handle("/api/v1/users", rl.Handle(v1.GetUserByLoginHandler(srv))).Methods(http.MethodGet)
	r.Handle("/api/v1/users/{id:[0-9]+}", rl.Handle(v1.GetUserByIDHandler(srv))).Methods(http.MethodGet)
	r.Handle("/api/v1/users", rl.Handle(v1.RegisterUserHandler(srv))).Methods(http.MethodPost)

	// expense service - v1
	r.Handle("/api/v1/expenses", rl.Handle(am.Authenticate(v1.UpdateExpenseHandler(srv)))).Methods(http.MethodPut, http.MethodPost)
	r.Handle("/api/v1/categories", rl.Handle(am.Authenticate(v1.UpdateCategoryHandler(srv)))).Methods(http.MethodPut, http.MethodPost)

	srv.Log().InfoWithFields(logger.Fields{
		"addr": srv.Info().Address.String(),
	}, "Service initialized. Listening for incomming connections")

	discovery.Serve(discoveryAddr)

	http.ListenAndServe(srv.Info().Address.String(), r)
}
