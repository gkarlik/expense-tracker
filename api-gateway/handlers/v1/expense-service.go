package v1

import (
	"net/http"

	"github.com/gkarlik/expense-tracker/api-gateway/handlers"
	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/quark-go"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetExpenseServiceConn(s quark.Service) (*grpc.ClientConn, error) {
	url, err := s.Discovery().GetServiceAddress(sd.ByName("ExpenseService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		handlers.LogError(s, err, "Cannot locate ExpenseService")
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		handlers.LogError(s, err, "Cannot dial address provided for ExpenseService")
		return nil, err
	}
	return conn, err
}

func UpdateExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense es.ExpenseRequest

		if err := handlers.ParseRequestData(s, r, &expense); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		e, err := expenseService.UpdateExpense(context.Background(), &expense)
		if err != nil {
			handlers.LogError(s, err, "Cannot update expense")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := handlers.JSONResponse(s, w, e); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category es.CategoryRequest

		if err := handlers.ParseRequestData(s, r, &category); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		c, err := expenseService.UpdateCategory(context.Background(), &category)
		if err != nil {
			handlers.LogError(s, err, "Cannot update expense")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := handlers.JSONResponse(s, w, c); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
