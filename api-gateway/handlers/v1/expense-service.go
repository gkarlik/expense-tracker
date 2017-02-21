package v1

import (
	"net/http"

	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetExpenseServiceConn(s quark.Service) (*grpc.ClientConn, error) {
	url, err := s.Discovery().GetServiceAddress(sd.ByName("ExpenseService"), sd.ByVersion("v1"))
	if err != nil || url == nil {
		return nil, err
	}
	conn, err := grpc.Dial(url.String(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, err
}

func UpdateExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense es.ExpenseRequest

		if err := handler.ParseRequestData(r, &expense); err != nil {
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}
		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		e, err := expenseService.UpdateExpense(context.Background(), &expense)
		if err != nil {
			handler.ErrorResponse(w, errors.ErrCannotUpdateExpense, http.StatusInternalServerError)
			return
		}
		handler.Response(w, e)
	}
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category es.CategoryRequest

		if err := handler.ParseRequestData(r, &category); err != nil {
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}
		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		c, err := expenseService.UpdateCategory(context.Background(), &category)
		if err != nil {
			handler.ErrorResponse(w, errors.ErrCannotUpdateCategory, http.StatusInternalServerError)
			return
		}
		handler.Response(w, c)
	}
}
