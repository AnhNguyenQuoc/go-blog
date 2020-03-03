package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	router := httprouter.New()
	router.GET("/", HelloWorld)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
