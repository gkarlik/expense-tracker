package main

import (
	"strconv"
	"time"

	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/user-service/v1/model"
	"github.com/gkarlik/expense-tracker/user-service/v1/proxy"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/circuitbreaker"
	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/metrics/prometheus"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/gkarlik/quark-go/service/discovery/plain"
	gRPC "github.com/gkarlik/quark-go/service/rpc/grpc"
	nt "github.com/gkarlik/quark-go/service/trace/noop"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type UserService struct {
	*quark.ServiceBase
}

func CreateUserService() *UserService {
	name := quark.GetEnvVar("USER_SERVICE_NAME")
	version := quark.GetEnvVar("USER_SERVICE_VERSION")
	p := quark.GetEnvVar("USER_SERVICE_PORT")
	discovery := quark.GetEnvVar("DISCOVERY")

	port, err := strconv.Atoi(p)
	if err != nil {
		panic("Incorrect port value!")
	}

	addr, err := quark.GetHostAddress(port)
	if err != nil {
		panic("Cannot resolve host address!")
	}

	return &UserService{
		ServiceBase: quark.NewService(
			quark.Name(name),
			quark.Version(version),
			quark.Address(addr),
			quark.Discovery(plain.NewServiceDiscovery(discovery)),
			quark.Metrics(prometheus.NewMetricsExposer()),
			quark.Tracer(nt.NewTracer())),
	}
}

var (
	dialect   = quark.GetEnvVar("DATABASE_DIALECT")
	dbConnStr = quark.GetEnvVar("DATABASE")
	service   = CreateUserService()
)

func NewDbContext() (rdbms.DbContext, error) {
	context, err := gorm.NewDbContext(dialect, dbConnStr)
	if err != nil {
		return nil, err
	}

	// set database table name format
	context.DB.SingularTable(true)

	return context, nil
}

func UpgradeDatabase(s quark.Service) error {
	if quark.GetEnvVar("UPGRADE_DATABASE") != "1" {
		s.Log().Info("Database upgrade is disabled")

		return nil
	}

	s.Log().Info("Upgrading database ...")

	context, err := NewDbContext()
	if err != nil {
		s.Log().ErrorWithFields(logger.Fields{"error": err}, "Cannot create database context")
		return err
	}

	context.(*gorm.DbContext).DB.DropTable(&model.User{})
	context.(*gorm.DbContext).DB.AutoMigrate(&model.User{})

	s.Log().Info("Database upgrade completed")

	return nil
}

func (us *UserService) RegisterUser(ctx context.Context, in *proxy.UserRequest) (*proxy.RegisterUserResponse, error) {
	span := quark.StartRPCSpan(ctx, service, "user_service_register_user")
	defer span.Finish()

	if (proxy.UserRequest{}) == *in {
		return nil, errors.ErrInvalidUserModel
	}

	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewUserRepository(context)

	// check if user already exists
	u, _ := repo.FindByLogin(in.Login)
	if u != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(in.Pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Login:     in.Login,
		ID:        uuid.NewV4().String(),
		Password:  string(hashedPassword),
		Pin:       string(hashedPin),
	}

	if ok := user.IsValid(); !ok {
		return nil, errors.ErrInvalidUserModel
	}

	if err := repo.Save(user); err != nil {
		return nil, err
	}

	return &proxy.RegisterUserResponse{
		VerificationLink: "TODO: send email and generate link",
	}, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, in *proxy.UserCredentialsRequest) (*proxy.UserIDResponse, error) {
	span := quark.StartRPCSpan(ctx, service, "user_service_authenticate_user")
	defer span.Finish()

	if in.Login == "" {
		return nil, errors.ErrInvalidRequestData
	}

	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewUserRepository(context)
	user, err := repo.FindByLogin(in.Login)
	if err != nil {
		return nil, err
	}

	if in.Pin == "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
			return nil, err
		}
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(in.Pin)); err != nil {
			return nil, err
		}
	}
	return &proxy.UserIDResponse{
		ID: user.ID,
	}, nil
}

func (us *UserService) GetUserByID(ctx context.Context, in *proxy.UserIDRequest) (*proxy.UserResponse, error) {
	span := quark.StartRPCSpan(ctx, service, "user_service_get_user_by_id")
	defer span.Finish()

	if in.ID == "" {
		return nil, errors.ErrInvalidRequestParameters
	}

	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewUserRepository(context)
	user, err := repo.FindByID(in.ID)
	if err != nil {
		return nil, err
	}

	return &proxy.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Login:     user.Login,
	}, nil
}

func (us *UserService) GetUserByLogin(ctx context.Context, in *proxy.UserLoginRequest) (*proxy.UserResponse, error) {
	span := quark.StartRPCSpan(ctx, service, "user_service_get_user_by_login")
	defer span.Finish()

	if in.Login == "" {
		return nil, errors.ErrInvalidRequestParameters
	}

	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewUserRepository(context)
	user, err := repo.FindByLogin(in.Login)
	if err != nil {
		return nil, err
	}

	return &proxy.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Login:     user.Login,
	}, nil
}

func (us *UserService) UpdateUser(ctx context.Context, in *proxy.UpdateUserRequest) (*proxy.UserResponse, error) {
	span := quark.StartRPCSpan(ctx, service, "user_service_update_user")
	defer span.Finish()

	if (proxy.UpdateUserRequest{}) == *in {
		return nil, errors.ErrInvalidUserModel
	}

	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(in.Pin), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Login:     in.Login,
		Password:  string(hashedPassword),
		Pin:       string(hashedPin),
	}

	if ok := user.IsValid(); !ok {
		return nil, errors.ErrInvalidUserModel
	}

	repo := model.NewUserRepository(context)
	if err := repo.Save(user); err != nil {
		return nil, err
	}

	return &proxy.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Login:     user.Login,
	}, nil
}

func (us *UserService) RegisterServiceInstance(server interface{}, serviceInstance interface{}) error {
	proxy.RegisterUserServiceServer(server.(*grpc.Server), serviceInstance.(proxy.UserServiceServer))

	return nil
}

func main() {
	cb := &circuitbreaker.DefaultCircuitBreaker{}

	_, err := cb.Execute(func() (interface{}, error) {
		return nil, UpgradeDatabase(service)
	}, circuitbreaker.Retry(3), circuitbreaker.Timeout(5*time.Second))

	if err != nil {
		panic("Cannot upgrade database!")
	}

	_, err = cb.Execute(func() (interface{}, error) {
		return nil, service.Discovery().RegisterService(sd.WithInfo(service.Info()))
	}, circuitbreaker.Retry(3), circuitbreaker.Timeout(10*time.Second))

	if err != nil {
		service.Log().ErrorWithFields(logger.Fields{
			"err": err,
		}, "Cannot register service")

		panic("Cannot register service!")
	}

	done := quark.HandleInterrupt(service)
	server := gRPC.NewServer()
	defer func() {
		server.Dispose()
		service.Dispose()
	}()

	go func() {
		service.Metrics().Expose()
	}()

	go func() {
		server.Start(service)
	}()

	<-done
}
