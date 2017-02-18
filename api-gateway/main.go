package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	us "github.com/gkarlik/expense-tracker/api-gateway/proxy/user-service/v1"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/metrics/noop"
	"github.com/gkarlik/quark-go/ratelimiter"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/gkarlik/quark-go/service/discovery/plain"
	nt "github.com/gkarlik/quark-go/service/trace/noop"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	componentName = "ApiGateway"
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

func getUserServiceConn() (*grpc.ClientConn, error) {
	url, err := srv.Discovery().GetServiceAddress(sd.ByName("UserService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot locate UserService")
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
			"address":   url,
		}, "Cannot dial address provided for UserService")
		return nil, err
	}
	return conn, err
}

func getExpenseServiceConn() (*grpc.ClientConn, error) {
	url, err := srv.Discovery().GetServiceAddress(sd.ByName("ExpenseService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot locate ExpenseService")
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
			"address":   url,
		}, "Cannot dial address provided for ExpenseService")
		return nil, err
	}
	return conn, err
}

func main() {
	defer srv.Dispose()

	secret := quark.GetEnvVar("GATEWAY_SECRET")
	am := auth.NewAuthenticationMiddleware(
		auth.WithSecret(secret),
		auth.WithContextKey("USER_KEY"),
		auth.WithAuthenticationFunc(func(credentials auth.Credentials) (auth.Claims, error) {
			conn, err := getUserServiceConn()
			if err != nil || conn == nil {
				srv.Log().ErrorWithFields(logger.Fields{
					"error":     err,
					"component": componentName,
				}, "Cannot connect to UserService")
				return auth.Claims{}, errors.New("Cannot connect to UserService")
			}
			defer conn.Close()

			userService := us.NewUserServiceClient(conn)
			_, err = userService.AuthenticateUser(context.Background(), &us.AuthenticateUserRequest{
				Login:    credentials.Username,
				Password: credentials.Password,
				Pin:      credentials.Properties["Pin"],
			})
			if err != nil {
				srv.Log().ErrorWithFields(logger.Fields{
					"error": err,
				}, "Cannot authenticate user")
				return auth.Claims{}, errors.New("Invalid username or password")
			}
			return auth.Claims{
				Username: credentials.Username,
				StandardClaims: jwt.StandardClaims{
					Issuer:    srv.Info().Address.String(),
					ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
				},
			}, nil
		}))

	rl := ratelimiter.NewHTTPRateLimiter(100 * time.Millisecond)

	r := mux.NewRouter()
	// user service
	r.Handle("/login", rl.Handle(http.HandlerFunc(am.GenerateToken))).Methods(http.MethodPost)
	r.Handle("/register", rl.Handle(http.HandlerFunc(registerUserHandler))).Methods(http.MethodPost)

	// expense service
	r.Handle("/expense/update", rl.Handle(http.HandlerFunc(updateExpenseHandler))).Methods(http.MethodPost)
	r.Handle("/category/update", rl.Handle(http.HandlerFunc(updateCategoryHandler))).Methods(http.MethodPost)

	srv.Log().InfoWithFields(logger.Fields{
		"addr": srv.Info().Address.String(),
	}, "Service initialized. Listening for incomming connections")

	discovery.Serve(discoveryAddr)

	http.ListenAndServe(srv.Info().Address.String(), r)
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var user us.RegisterUserRequest

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &user)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot process user registration request")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	conn, err := getUserServiceConn()
	if err != nil || conn == nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot connect to UserService")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	userService := us.NewUserServiceClient(conn)
	_, err = userService.RegisterUser(context.Background(), &user)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot register user")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var expense es.ExpenseRequest

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &expense)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot process expense update request")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	conn, err := getExpenseServiceConn()
	if err != nil || conn == nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot connect to ExpenseService")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	expenseService := es.NewExpenseServiceClient(conn)
	e, err := expenseService.UpdateExpense(context.Background(), &expense)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot update expense")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(e)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot send JSON reponse")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category es.CategoryRequest

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &category)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot process expense update request")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	conn, err := getExpenseServiceConn()
	if err != nil || conn == nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot connect to ExpenseService")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	expenseService := es.NewExpenseServiceClient(conn)
	c, err := expenseService.UpdateCategory(context.Background(), &category)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot update expense")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(c)
	if err != nil {
		srv.Log().ErrorWithFields(logger.Fields{
			"error":     err,
			"component": componentName,
		}, "Cannot send JSON reponse")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
