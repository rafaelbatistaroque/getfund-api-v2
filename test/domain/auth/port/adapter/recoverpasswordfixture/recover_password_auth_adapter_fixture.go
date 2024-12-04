package recoverpasswordfixture

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(GetRecoverPasswordInputSerialized())
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	}
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func GetRecoverPasswordInputSerialized() string {
	return `{"username": "fake-username"}`
}
