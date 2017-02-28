package v1

import (
	"net/http"
	"strconv"
	"time"

	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
	sd "github.com/gkarlik/quark-go/service/discovery"
	"github.com/gorilla/mux"
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

func validateExpenseRequest(er *es.ExpenseRequest, r *http.Request) error {
	if er.ID == "" && r.Method == http.MethodPut {
		return errors.ErrInvalidRequestParameters
	}
	if r.Method == http.MethodPost {
		er.ID = uuid.NewV4().String()
	}
	er.UserID = handler.GetUserID(r)

	if er.CategoryID == "" || (er.Date <= 0 && er.Date > time.Now().Unix()) || er.Value <= 0 {
		return errors.ErrInvalidRequestParameters
	}
	return nil
}

func UpdateExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetRequestID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update expense handler")

		var expense es.ExpenseRequest
		body, err := handler.ParseRequestData(r, &expense)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusBadRequest)
			return
		}

		err = validateExpenseRequest(&expense, r)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Invalid request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusBadRequest)
			return
		}

		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "body": string(body)}, "Update expense request body")

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
		handler.Response(w, e, http.StatusOK)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing update expense handler")
	}
}

func GetExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetRequestID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get expense handler")

		id := mux.Vars(r)["id"]
		if id == "" {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID}, "Missing 'ID' parameter in request")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to ExpenseService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		e, err := expenseService.GetExpense(context.Background(), &es.ExpenseIDRequest{ID: id})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get expense by ID")
			handler.ErrorResponse(w, errors.ErrExpenseNotFound, http.StatusNotFound)
			return
		}
		handler.Response(w, e, http.StatusOK)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get expense handler")
	}
}

func RemoveExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetRequestID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing remove expense handler")

		id := mux.Vars(r)["id"]
		if id == "" {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID}, "Missing 'ID' parameter in request")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to ExpenseService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		_, err = expenseService.RemoveExpense(context.Background(), &es.ExpenseIDRequest{ID: id})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot remove expense")
			handler.ErrorResponse(w, errors.ErrCannotRemoveExpense, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing remove expense handler")
	}
}

func GetExpensesHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetRequestID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get expenses handler")

		q := r.URL.Query()
		p := q.Get("p")
		if p == "" {
			// is 'p' parameter is missing default it to '0'
			p = "0"
		}

		page, err := strconv.Atoi(p)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID}, "Invalid 'page' parameter in request")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		conn, err := GetExpenseServiceConn(s)
		if err != nil || conn == nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot connect to ExpenseService")
			handler.ErrorResponse(w, errors.ErrInternal, http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		expenseService := es.NewExpenseServiceClient(conn)
		es, err := expenseService.GetUserExpenses(context.Background(), &es.UserPagingRequest{
			Limit:  handler.ExpensesPageSize,
			Offset: int32(page) * handler.ExpensesPageSize,
			UserID: handler.GetUserID(r),
		})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get expenses")
			handler.ErrorResponse(w, errors.ErrCannotGetExpenses, http.StatusInternalServerError)
			return
		}
		if len(es.Expenses) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			handler.Response(w, es, http.StatusOK)
		}

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get expenses handler")
	}
}