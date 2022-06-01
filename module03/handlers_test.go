package module03

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestHealthz(t *testing.T) {
	route := Route{
		Name:        "healthz",
		Method:      "GET",
		Path:        "/healthz",
		HandlerFunc: healthz,
	}
	req, err := http.NewRequest(route.Method, route.Path, nil)
	if err != nil {
		log.Fatal(err)
	}
	recorder := newHttpRecorder(req, route)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected response code to be 200, actual %d\n", recorder.Code)
	}
	expected := "Hello LiveRamp SRE"
	if recorder.Body.String() != expected {
		t.Errorf("Expected response body %s actual %s\n", expected, recorder.Body.String())
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newHttpRecorder(req *http.Request, r Route) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(r.Method, r.Path, r.HandlerFunc)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}
