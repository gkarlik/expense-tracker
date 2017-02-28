package v1

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	us "github.com/gkarlik/expense-tracker/api-gateway/proxy/user-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
	"github.com/gkarlik/quark-go/logger"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetUserServiceConn(s quark.Service) (*grpc.ClientConn, error) {
	url, err := s.Discovery().GetServiceAddress(sd.ByName("UserService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func RegisterUserHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value("Request-ID")
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing register user handler")

		var user us.UserRequest
		body, err := handler.ParseRequestData(r, &user)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusBadRequest)
			return
		}
		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "body": string(body)}, "Register user request body")

		conn, err := GetUserServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to UserService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		userService := us.NewUserServiceClient(conn)
		_, err = userService.RegisterUser(context.Background(), &user)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot register user")
			if errors.ErrUserAlreadyExists.IsSame(err) {
				handler.ErrorResponse(w, errors.ErrUserAlreadyExists, http.StatusConflict)
				return
			}
			handler.ErrorResponse(w, errors.ErrCannotRegisterUser, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing register user handler")
	}
}

func AuthenticateUser(s quark.Service, credentials auth.Credentials) (auth.Claims, error) {
	s.Log().Info("Processing authenticate user handler")

	conn, err := GetUserServiceConn(s)
	if err != nil || conn == nil {
		s.Log().ErrorWithFields(logger.Fields{"error": err}, "Cannot connect to UserService")
		return auth.Claims{}, errors.ErrInternal
	}
	defer conn.Close()

	userService := us.NewUserServiceClient(conn)
	ur, err := userService.AuthenticateUser(context.Background(), &us.UserCredentialsRequest{
		Login:    credentials.Username,
		Password: credentials.Password,
		Pin:      credentials.Properties["Pin"],
	})
	if err != nil {
		s.Log().ErrorWithFields(logger.Fields{"error": err}, "Cannot authenticate user")
		return auth.Claims{}, errors.ErrInvalidUsernamePassword
	}
	s.Log().Info("Done processing authenticate user handler")

	return auth.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.Info().Address.String(),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		Properties: map[string]interface{}{"UserID": ur.ID},
	}, nil
}

func GetUserByLoginHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value("Request-ID")
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get user by login handler")

		q := r.URL.Query()
		login := q.Get("l")

		if login == "" {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID}, "Missing 'login' parameter in request")
			handler.ErrorResponse(w, errors.ErrInvalidRequestParameters, http.StatusBadRequest)
			return
		}

		conn, err := GetUserServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to UserService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		userService := us.NewUserServiceClient(conn)
		user, err := userService.GetUserByLogin(context.Background(), &us.UserLoginRequest{Login: login})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get user by login")
			handler.ErrorResponse(w, errors.ErrUserNotFound, http.StatusNotFound)
			return
		}
		handler.Response(w, user)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get user by login handler")
	}
}

func GetUserByIDHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value("Request-ID")
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get user by ID handler")

		userID := mux.Vars(r)["id"]

		if userID == "" {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID}, "Missing 'ID' parameter in request")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		conn, err := GetUserServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to UserService")
			handler.ErrorResponse(w, errors.ErrInvalidRequestParameters, http.StatusBadRequest)
			return
		}
		defer conn.Close()

		userService := us.NewUserServiceClient(conn)
		user, err := userService.GetUserByID(context.Background(), &us.UserIDRequest{ID: userID})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get user by ID")
			handler.ErrorResponse(w, errors.ErrUserNotFound, http.StatusNotFound)
			return
		}
		handler.Response(w, user)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get user by ID handler")
	}
}
