package v1

import (
	"net/http"
	"strconv"

	es "github.com/gkarlik/expense-tracker/api-gateway/proxy/expense-service/v1"
	"github.com/gkarlik/expense-tracker/shared/errors"
	"github.com/gkarlik/expense-tracker/shared/handler"
	"github.com/gkarlik/quark-go"
	"github.com/gkarlik/quark-go/logger"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

func validateCategoryRequest(cr *es.CategoryRequest, r *http.Request) error {
	if cr.ID == "" && r.Method == http.MethodPut {
		return errors.ErrInvalidRequestParameters
	}
	if r.Method == http.MethodPost {
		cr.ID = uuid.NewV4().String()
	}
	cr.UserID = handler.GetUserID(r)

	if cr.Name == "" || cr.Limit <= 0 {
		return errors.ErrInvalidRequestParameters
	}
	return nil
}

func UpdateCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetRequestID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing update category handler")

		var category es.CategoryRequest
		body, err := handler.ParseRequestData(r, &category)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot parse request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusBadRequest)
			return
		}
		err = validateCategoryRequest(&category, r)
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Invalid request data")
			handler.ErrorResponse(w, errors.ErrInvalidRequestData, http.StatusBadRequest)
			return
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
		handler.Response(w, c, http.StatusOK)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing update category handler")
	}
}

func GetCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetUserID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get category handler")

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
		c, err := expenseService.GetCategory(context.Background(), &es.CategoryIDRequest{ID: id})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get category by ID")
			handler.ErrorResponse(w, errors.ErrCategoryNotFound, http.StatusNotFound)
			return
		}
		handler.Response(w, c, http.StatusOK)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get category handler")
	}
}

func RemoveCategoryHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetUserID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing remove category handler")

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
		_, err = expenseService.RemoveCategory(context.Background(), &es.CategoryIDRequest{ID: id})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot remove category")
			handler.ErrorResponse(w, errors.ErrCannotRemoveCategory, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get category handler")
	}
}

func GetCategoriesHandler(s quark.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := handler.GetUserID(r)
		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Processing get categories handler")

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
		cs, err := expenseService.GetUserCategories(context.Background(), &es.UserPagingRequest{
			Limit:  handler.CategoriesPageSize,
			Offset: int32(page) * handler.CategoriesPageSize,
			UserID: handler.GetUserID(r),
		})
		if err != nil {
			s.Log().ErrorWithFields(logger.Fields{"requestID": reqID, "error": err}, "Cannot get categories")
			handler.ErrorResponse(w, errors.ErrCannotGetCategories, http.StatusInternalServerError)
			return
		}
		if len(cs.Categories) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			handler.Response(w, cs, http.StatusOK)
		}

		s.Log().InfoWithFields(logger.Fields{"requestID": reqID}, "Done processing get categories handler")
	}
}
