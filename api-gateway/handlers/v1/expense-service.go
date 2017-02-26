package v1

import (
	"net/http"
	"time"

	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	auth "github.com/gkarlik/quark-go/auth/jwt"
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

func validateExpenseRequest(er *es.ExpenseRequest, r *http.Request) error {
	if er.ID == "" {
		er.ID = uuid.NewV4().String()
	}
	claims := r.Context().Value("USER_KEY").(auth.Claims)
	userID := claims.Properties["UserID"].(string)
	er.UserID = userID

	if er.CategoryID == "" || (er.Date <= 0 && er.Date > time.Now().Unix()) || er.Value <= 0 {
		return errors.ErrInvalidRequestParameters
	}
	return nil
}

func UpdateExpenseHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value("Request-ID")
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update expense handler")

		var expense es.ExpenseRequest
		body, err := handler.ParseRequestData(r, &expense)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}

		err = validateExpenseRequest(&expense, r)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Invalid request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
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
		handler.Response(w, e)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing update expense handler")
	}
}

func validateCategoryRequest(cr *es.CategoryRequest, r *http.Request) error {
	if cr.ID == "" {
		cr.ID = uuid.NewV4().String()
	}
	claims := r.Context().Value("USER_KEY").(auth.Claims)
	userID := claims.Properties["UserID"].(string)
	cr.UserID = userID

	if cr.Name == "" || cr.Limit <= 0 {
		return errors.ErrInvalidRequestParameters
	}
	return nil
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value("Request-ID")
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update category handler")

		var category es.CategoryRequest
		body, err := handler.ParseRequestData(r, &category)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
			return
		}
		err = validateCategoryRequest(&category, r)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Invalid request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusInternalServerError)
		}

		s.Log().DebugWithFields(logger.Fields{"requestID": reqID, "body": string(body)}, "Update category request body")

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
