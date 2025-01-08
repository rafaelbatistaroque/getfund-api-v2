package response_proxy

import (
	"encoding/json"
	"getfund-api-v2/internal/shared/result_app"
	"net/http"
)

var httpCodeMap = map[int]int{
	result_app.SUCCESS_CODE:             http.StatusOK,                  // 200
	result_app.UNAUTHORIZED_CODE:        http.StatusUnauthorized,        // 401
	result_app.DUPLICATED_ENTRY_CODE:    http.StatusConflict,            // 409
	result_app.NOT_FOUND_CODE:           http.StatusNotFound,            // 404
	result_app.SERVER_ERROR_CODE:        http.StatusInternalServerError, // 500
	result_app.CONSTRAINT_VIOLATED_CODE: http.StatusUnprocessableEntity, // 422
	result_app.BAD_REQUEST_CODE:         http.StatusBadRequest,          // 400
	result_app.UNAVAILABLE_CODE:         http.StatusServiceUnavailable,  // 503
	result_app.UNMODIFIED_CODE:          http.StatusNotModified,         // 304
	result_app.SUCCESS_CREATED_CODE:     http.StatusCreated,             // 201
}

type handleFunc func(http.ResponseWriter, *http.Request) (interface{}, int, error)

type ResponseBody struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

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

func SetError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(getHTTPCode(code))
	json.NewEncoder(w).Encode(ResponseBody{
		Code: code,
		Data: err.Error()})
}

func SetSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(getHTTPCode(code))
	json.NewEncoder(w).Encode(ResponseBody{
		Code: code,
		Data: data})
}

func getHTTPCode(appCode int) int {
	if httpCode, exists := httpCodeMap[appCode]; exists {
		return httpCode
	}
	return http.StatusInternalServerError
}
