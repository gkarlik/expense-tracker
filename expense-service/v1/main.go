package main

import (
	"strconv"

	"github.com/gkarlik/expense-tracker/expense-service/v1/model"
	"github.com/gkarlik/expense-tracker/expense-service/v1/proxy"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/data/access/rdbms"
	"github.com/gkarlik/quark-go/data/access/rdbms/gorm"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gkarlik/quark-go/metrics/noop"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/gkarlik/quark-go/service/discovery/plain"
	gRPC "github.com/gkarlik/quark-go/service/rpc/grpc"
	nt "github.com/gkarlik/quark-go/service/trace/noop"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ExpenseService struct {
	*quark.ServiceBase
}

func CreateExpenseService() *ExpenseService {
	name := quark.GetEnvVar("EXPENSE_SERVICE_NAME")
	version := quark.GetEnvVar("EXPENSE_SERVICE_VERSION")
	p := quark.GetEnvVar("EXPENSE_SERVICE_PORT")
	discovery := quark.GetEnvVar("DISCOVERY")

	port, err := strconv.Atoi(p)
	if err != nil {
		panic("Incorrect port value!")
	}

	addr, err := quark.GetHostAddress(port)
	if err != nil {
		panic("Cannot resolve host address!")
	}

	return &ExpenseService{
		ServiceBase: quark.NewService(
			quark.Name(name),
			quark.Version(version),
			quark.Address(addr),
			quark.Discovery(plain.NewServiceDiscovery(discovery)),
			quark.Metrics(noop.NewMetricsReporter()),
			quark.Tracer(nt.NewTracer())),
	}
}

var (
	dialect   = quark.GetEnvVar("DATABASE_DIALECT")
	dbConnStr = quark.GetEnvVar("DATABASE")
	service   = CreateExpenseService()
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

	context.(*gorm.DbContext).DB.DropTable(&model.Category{})
	context.(*gorm.DbContext).DB.DropTable(&model.Expense{})
	context.(*gorm.DbContext).DB.AutoMigrate(&model.Category{})
	context.(*gorm.DbContext).DB.AutoMigrate(&model.Expense{})

	s.Log().Info("Database upgrade completed")

	return nil
}

func (es *ExpenseService) GetExpense(ctx context.Context, in *proxy.ExpenseIDRequest) (*proxy.ExpenseResponse, error) {
	return nil, nil
}

func (es *ExpenseService) UpdateExpense(ctx context.Context, in *proxy.ExpenseRequest) (*proxy.ExpenseResponse, error) {
	return nil, nil
}

func (es *ExpenseService) RemoveExpense(ctx context.Context, in *proxy.ExpenseIDRequest) (*proxy.DeleteExpenseResponse, error) {
	return nil, nil
}

func (es *ExpenseService) GetUserExpenses(ctx context.Context, in *proxy.UserPagingRequest) (*proxy.ExpensesReponse, error) {
	return nil, nil
}

func (es *ExpenseService) GetCategory(ctx context.Context, in *proxy.CategoryIDRequest) (*proxy.CategoryResponse, error) {
	return nil, nil
}

func (es *ExpenseService) UpdateCategory(ctx context.Context, in *proxy.CategoryRequest) (*proxy.CategoryResponse, error) {
	return nil, nil
}

func (es *ExpenseService) RemoveCategory(ctx context.Context, in *proxy.CategoryIDRequest) (*proxy.DeleteCategoryResponse, error) {
	return nil, nil
}

func (es *ExpenseService) GetUserCategories(ctx context.Context, in *proxy.UserPagingRequest) (*proxy.CategoriesReponse, error) {
	return nil, nil
}

func (es *ExpenseService) RegisterServiceInstance(server interface{}, serviceInstance interface{}) error {
	proxy.RegisterExpenseServiceServer(server.(*grpc.Server), serviceInstance.(proxy.ExpenseServiceServer))

	return nil
}

func main() {
	if err := UpgradeDatabase(service); err != nil {
		panic("Cannot upgrade database!")
	}

	if err := service.Discovery().RegisterService(sd.WithInfo(service.Info())); err != nil {
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
		server.Start(service)
	}()

	<-done
}
