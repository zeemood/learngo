package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errNotExist(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	panic("Unknown Error")
}
func noError(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return userError(`unknown user error`)
}

type userError string

func (err userError) Error() string {
	return err.Message()
}

func (err userError) Message() string {
	return string(err)
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errNotExist, http.StatusNotFound, `Not Found`},
	{errUnknown, http.StatusInternalServerError, `Internal Server Error`},
	{noError, 200, ""},
	{errUserError, 400, "unknown user error"},
}

func TestErrorWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errorWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}

}

func TestHttpServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, code int, message string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.TrimSpace(string(b))
	if body != message || resp.StatusCode != code {
		t.Errorf("expected (%d,%s) but get(%d,%s)", code, message, resp.StatusCode, body)
	}
}
