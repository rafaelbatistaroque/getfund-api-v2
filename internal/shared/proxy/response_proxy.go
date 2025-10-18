package shared_response_proxy

import (
	"encoding/json"
	shared_error "getfund-api-v2/internal/shared/error"

	"net/http"
)

var httpCodeMap = map[int]int{
	shared_error.SUCCESS_CODE:               http.StatusOK,                  // 200
	shared_error.UNAUTHORIZED_CODE:          http.StatusUnauthorized,        // 401
	shared_error.DUPLICATED_ENTRY_CODE:      http.StatusConflict,            // 409
	shared_error.NOT_FOUND_CODE:             http.StatusNotFound,            // 404
	shared_error.UNPROCESSABLE_CONTENT_CODE: http.StatusUnprocessableEntity, // 422
	shared_error.SERVER_ERROR_CODE:          http.StatusInternalServerError, // 500
	shared_error.CONSTRAINT_VIOLATED_CODE:   http.StatusUnprocessableEntity, // 422
	shared_error.BAD_REQUEST_CODE:           http.StatusBadRequest,          // 400
	shared_error.UNAVAILABLE_CODE:           http.StatusServiceUnavailable,  // 503
	shared_error.UNMODIFIED_CODE:            http.StatusNotModified,         // 304
	shared_error.SUCCESS_CREATED_CODE:       http.StatusCreated,             // 201
}

type handleFunc func(http.ResponseWriter, *http.Request) (any, int, error)

// ResponseBody is the standardized JSON response structure.
type ResponseBody struct {
	Code  int    `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// New wraps a custom handler function and returns a standard http.HandlerFunc.
// It orchestrates the response standardization, converting the return values
// of the handle func into a consistent JSON response.
func New(handle handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result, code, err := handle(w, r)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			SetError(w, code, err)
			return
		}

		SetSuccess(w, code, result)
	}
}

// SetError writes a standardized error response to the http.ResponseWriter.
func SetError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(getHTTPCode(code))
	json.NewEncoder(w).Encode(ResponseBody{
		Code:  code,
		Error: err.Error(),
	})
}

// SetSuccess writes a standardized success response to the http.ResponseWriter.
func SetSuccess(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(getHTTPCode(code))
	json.NewEncoder(w).Encode(ResponseBody{
		Code: code,
		Data: data,
	})
}

func getHTTPCode(appCode int) int {
	if httpCode, exists := httpCodeMap[appCode]; exists {
		return httpCode
	}
	return http.StatusInternalServerError
}
