package v1

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gkarlik/expense-tracker/api-gateway/handlers"
	us "github.com/gkarlik/expense-tracker/api-gateway/proxy/user-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/logger"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetUserServiceConn(s quark.Service) (*grpc.ClientConn, error) {
	url, err := s.Discovery().GetServiceAddress(sd.ByName("UserService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		s.Log().ErrorWithFields(logger.Fields{"error": err}, "Cannot locate UserService")
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		s.Log().ErrorWithFields(logger.Fields{"error": err, "address": url}, "Cannot dial address provided for UserService")
		return nil, err
	}
	return conn, err
}

func RegisterUserHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user us.RegisterUserRequest

		if err := handlers.ParseRequestData(r, &user); err != nil {
			handlers.LogError(s, err, "Cannot process user registration request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		conn, err := GetUserServiceConn(s)
		if err != nil || conn == nil {
			handlers.LogError(s, err, "Cannot connect to UserService")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		userService := us.NewUserServiceClient(conn)
		_, err = userService.RegisterUser(context.Background(), &user)
		if err != nil {
			handlers.LogError(s, err, "Cannot register user")
			if errors.ErrUserExists.IsSame(err) {
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func AuthenticateUser(s quark.Service, credentials auth.Credentials) (auth.Claims, error) {
	conn, err := GetUserServiceConn(s)
	if err != nil || conn == nil {
		handlers.LogError(s, err, "Cannot connect to UserService")
		return auth.Claims{}, errors.ErrUserServiceConnection
	}
	defer conn.Close()

	userService := us.NewUserServiceClient(conn)
	_, err = userService.AuthenticateUser(context.Background(), &us.AuthenticateUserRequest{
		Login:    credentials.Username,
		Password: credentials.Password,
		Pin:      credentials.Properties["Pin"],
	})
	if err != nil {
		handlers.LogError(s, err, "Cannot authenticate user")
		return auth.Claims{}, errors.ErrInvalidUsernamePassword
	}
	return auth.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.Info().Address.String(),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}, nil
}
