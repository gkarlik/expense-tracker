package v1

import (
	"net/http"

	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/satori/go.uuid"
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
		reqID := uuid.NewV4()
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update expense handler")
		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "request": r}, "Update expense request")

		var expense es.ExpenseRequest
		body, err := handler.ParseRequestData(r, &expense)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}
		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "body": body}, "Update expense request body")

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to ExpenseService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		e, err := expenseService.UpdateExpense(context.Background(), &expense)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot update expense")
			handler.ErrorResponse(w, errors.ErrCannotUpdateExpense, http.StatusInternalServerError)
			return
		}
		handler.Response(w, e)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing update expense handler")
	}
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.NewV4()
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update category handler")
		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "request": r}, "Update category request")

		var category es.CategoryRequest
		body, err := handler.ParseRequestData(r, &category)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}
		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "body": body}, "Update category request body")

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to ExpenseService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		c, err := expenseService.UpdateCategory(context.Background(), &category)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot update category")
			handler.ErrorResponse(w, errors.ErrCannotUpdateCategory, http.StatusInternalServerError)
			return
		}
		handler.Response(w, c)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing update category handler")
	}
}
