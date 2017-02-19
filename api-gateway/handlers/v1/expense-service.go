package v1

import (
	"encoding/json"
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

		if err := handlers.ParseRequestData(r, &expense); err != nil {
			handlers.LogError(s, err, "Cannot process expense update request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			handlers.LogError(s, err, "Cannot connect to ExpenseService")
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

		data, err := json.Marshal(e)
		if err != nil {
			handlers.LogError(s, err, "Cannot send JSON reponse")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category es.CategoryRequest

		if err := handlers.ParseRequestData(r, &category); err != nil {
			handlers.LogError(s, err, "Cannot process expense update request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			handlers.LogError(s, err, "Cannot connect to ExpenseService")
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

		data, err := json.Marshal(c)
		if err != nil {
			handlers.LogError(s, err, "Cannot send JSON reponse")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
