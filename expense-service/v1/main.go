package main

import (
	"strconv"
	"time"

	"github.com/gkarlik/expense-tracker/expense-service/v1/model"
	"github.com/gkarlik/expense-tracker/expense-service/v1/proxy"
	"github.com/gkarlik/expense-tracker/shared/errors"
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
	uuid "github.com/satori/go.uuid"
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
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewExpenseRepository(context)
	expense, err := repo.FindByID(in.ID)
	if err != nil {
		return nil, err
	}

	return &proxy.ExpenseResponse{
		ID:         expense.ID,
		CategoryID: expense.CategoryID,
		Value:      expense.Value,
		Date:       expense.Date.Unix(),
	}, nil
}

func (es *ExpenseService) CreateExpense(ctx context.Context, in *proxy.CreateExpenseRequest) (*proxy.ExpenseResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	expense := &model.Expense{
		ID:         uuid.NewV4().String(),
		Date:       time.Unix(in.Date, 0),
		Value:      in.Value,
		UserID:     in.UserID,
		CategoryID: in.CategoryID,
	}

	if ok := expense.IsValid(); !ok {
		return nil, errors.ErrInvalidExpenseModel
	}

	repo := model.NewExpenseRepository(context)
	if err := repo.Save(expense); err != nil {
		return nil, err
	}

	return &proxy.ExpenseResponse{
		ID:         expense.ID,
		Date:       expense.Date.Unix(),
		Value:      expense.Value,
		CategoryID: expense.CategoryID,
	}, nil
}

func (es *ExpenseService) UpdateExpense(ctx context.Context, in *proxy.UpdateExpenseRequest) (*proxy.ExpenseResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	expense := &model.Expense{
		ID:         in.ID,
		Date:       time.Unix(in.Date, 0),
		Value:      in.Value,
		UserID:     in.UserID,
		CategoryID: in.CategoryID,
	}

	if ok := expense.IsValid(); !ok {
		return nil, errors.ErrInvalidExpenseModel
	}

	repo := model.NewExpenseRepository(context)
	if err := repo.Save(expense); err != nil {
		return nil, err
	}

	return &proxy.ExpenseResponse{
		ID:         expense.ID,
		Date:       expense.Date.Unix(),
		Value:      expense.Value,
		CategoryID: expense.CategoryID,
	}, nil
}

func (es *ExpenseService) RemoveExpense(ctx context.Context, in *proxy.ExpenseIDRequest) (*proxy.EmptyResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewExpenseRepository(context)
	expense, err := repo.FindByID(in.ID)
	if err != nil {
		return nil, err
	}
	if err := repo.Delete(expense); err != nil {
		return nil, err
	}
	return &proxy.EmptyResponse{}, nil
}

func (es *ExpenseService) GetUserExpenses(ctx context.Context, in *proxy.UserPagingRequest) (*proxy.ExpensesResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewExpenseRepository(context)
	expenses, err := repo.FindByUserID(in.UserID, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	var exps []*proxy.ExpenseResponse
	for i := 0; i < len(expenses); i++ {
		exp := &proxy.ExpenseResponse{
			ID:         expenses[i].ID,
			Date:       expenses[i].Date.Unix(),
			Value:      expenses[i].Value,
			CategoryID: expenses[i].CategoryID,
		}
		exps = append(exps, exp)
	}

	return &proxy.ExpensesResponse{
		Expenses: exps,
	}, nil
}

func (es *ExpenseService) GetCategory(ctx context.Context, in *proxy.CategoryIDRequest) (*proxy.CategoryResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewCategoryRepository(context)
	category, err := repo.FindByID(in.ID)
	if err != nil {
		return nil, err
	}
	return &proxy.CategoryResponse{
		ID:    category.ID,
		Limit: category.Limit,
		Name:  category.Name,
	}, nil
}

func (es *ExpenseService) CreateCategory(ctx context.Context, in *proxy.CreateCategoryRequest) (*proxy.CategoryResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	category := &model.Category{
		ID:     uuid.NewV4().String(),
		Limit:  in.Limit,
		Name:   in.Name,
		UserID: in.UserID,
	}

	if ok := category.IsValid(); !ok {
		return nil, errors.ErrInvalidCategoryModel
	}

	repo := model.NewExpenseRepository(context)
	if err := repo.Save(category); err != nil {
		return nil, err
	}

	return &proxy.CategoryResponse{
		ID:    category.ID,
		Limit: category.Limit,
		Name:  category.Name,
	}, nil
}

func (es *ExpenseService) UpdateCategory(ctx context.Context, in *proxy.UpdateCategoryRequest) (*proxy.CategoryResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	category := &model.Category{
		ID:     in.ID,
		Limit:  in.Limit,
		Name:   in.Name,
		UserID: in.UserID,
	}

	if ok := category.IsValid(); !ok {
		return nil, errors.ErrInvalidCategoryModel
	}

	repo := model.NewExpenseRepository(context)
	if err := repo.Save(category); err != nil {
		return nil, err
	}

	return &proxy.CategoryResponse{
		ID:    category.ID,
		Limit: category.Limit,
		Name:  category.Name,
	}, nil
}

func (es *ExpenseService) RemoveCategory(ctx context.Context, in *proxy.CategoryIDRequest) (*proxy.EmptyResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewCategoryRepository(context)
	category, err := repo.FindByID(in.ID)
	if err != nil {
		return nil, err
	}
	if err := repo.Delete(category); err != nil {
		return nil, err
	}
	return &proxy.EmptyResponse{}, nil
}

func (es *ExpenseService) GetUserCategories(ctx context.Context, in *proxy.UserPagingRequest) (*proxy.CategoriesResponse, error) {
	context, err := NewDbContext()
	if err != nil {
		return nil, err
	}
	defer context.Dispose()

	repo := model.NewCategoryRepository(context)
	categories, err := repo.FindByUserID(in.UserID, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	var cs []*proxy.CategoryResponse
	for i := 0; i < len(categories); i++ {
		c := &proxy.CategoryResponse{
			ID:    categories[i].ID,
			Name:  categories[i].Name,
			Limit: categories[i].Limit,
		}
		cs = append(cs, c)
	}

	return &proxy.CategoriesResponse{
		Categories: cs,
	}, nil
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
