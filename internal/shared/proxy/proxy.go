package proxy

import (
	"encoding/json"
	"getfund-api-v2/internal/shared/applicationcode"
	"net/http"
)

var httpCodeMap = map[int]int{
	applicationcode.CODE_SUCCESS:             http.StatusOK,                  // 200
	applicationcode.CODE_UNAUTHORIZED:        http.StatusUnauthorized,        // 401
	applicationcode.CODE_DUPLICATED_ENTRY:    http.StatusConflict,            // 409
	applicationcode.CODE_NOT_FOUND:           http.StatusNotFound,            // 404
	applicationcode.CODE_SERVER_ERROR:        http.StatusInternalServerError, // 500
	applicationcode.CODE_CONSTRAINT_VIOLATED: http.StatusBadRequest,          // 400
	applicationcode.BAD_REQUEST:              http.StatusBadRequest,          // 400
	applicationcode.CODE_UNAVAILABLE:         http.StatusServiceUnavailable,  // 503
	applicationcode.CODE_UNMODIFIED:          http.StatusNotModified,         // 304
}

type handleFunc func(http.ResponseWriter, *http.Request) (interface{}, int, error)

type ResponseBody struct {
	IsOk bool        `json:"isOk"`
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
		IsOk: false,
		Code: code,
		Data: err.Error()})
}

func SetSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(getHTTPCode(code))
	json.NewEncoder(w).Encode(ResponseBody{
		IsOk: true,
		Code: code,
		Data: data})
}

func getHTTPCode(appCode int) int {
	if httpCode, exists := httpCodeMap[appCode]; exists {
		return httpCode
	}
	return http.StatusInternalServerError // 500
}
