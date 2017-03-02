package main

import (
	"net/http"
	"strconv"

	v1 "github.com/gkarlik/expense-tracker/api-gateway/routes/v1"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/metrics/noop"
	"github.com/gkarlik/quark-go/service/discovery/plain"
	nt "github.com/gkarlik/quark-go/service/trace/noop"
	"github.com/gorilla/handlers"
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

	g := &Gateway{
		ServiceBase: quark.NewService(
			quark.Name(name),
			quark.Version(version),
			quark.Address(addr),
			quark.Discovery(discovery),
			quark.Metrics(noop.NewMetricsReporter()),
			quark.Tracer(nt.NewTracer())),
	}
	g.Log().SetLevel(logger.DebugLevel)

	return g
}

var srv = CreateGateway()

func main() {
	defer srv.Dispose()

	r := mux.NewRouter().StrictSlash(true)

	api := v1.Init(r, srv)
	v1.InitUsersRoutes(api, srv)
	v1.InitExpensesRoutes(api, srv)
	v1.InitCategoriesRoutes(api, srv)

	srv.Log().InfoWithFields(logger.Fields{
		"addr": srv.Info().Address.String(),
	}, "Service initialized. Listening for incomming connections")

	discovery.Serve(discoveryAddr)

	http.ListenAndServe(srv.Info().Address.String(), handlers.CORS(
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodPost, http.MethodPut, http.MethodDelete}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(r))
}
