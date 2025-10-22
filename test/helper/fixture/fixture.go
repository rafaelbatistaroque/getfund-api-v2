package fixture

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

type BaseFixture struct{}

func (f *BaseFixture) GetHttpRequestResponse(bodyString string) (w http.ResponseWriter, r *http.Request) {
	body := bytes.NewBufferString(bodyString)
	req := httptest.NewRequest("FAKE", "/", body)
	res := httptest.NewRecorder()

	return res, req
}

func (f *BaseFixture) GetHttpRequestResponseWithUrl(url string) (w http.ResponseWriter, r *http.Request) {
	req := httptest.NewRequest("FAKE", url, nil)
	res := httptest.NewRecorder()

	return res, req
}
