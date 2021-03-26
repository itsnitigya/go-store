package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/itsnitigya/go-store/app"
)

var a = &app.App{}

func TestMain(m *testing.M) {

	a.Initialize()
	code := m.Run()

	os.Exit(code)
}

func TestGetAndSet(t *testing.T) {

	response := executePostRequest([]byte(`{"key":"abc-1", "value": "1"}`))

	checkResponseCode(t, http.StatusCreated, response.Code)

	req, _ := http.NewRequest("GET", "/get/abc-1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		if strings.Compare(strings.Trim(body, "\""), "1") != 0 {
			t.Errorf("Expected 1. Got %s", body)
		}
	}
}

func TestPrefix(t *testing.T) {

	response := executePostRequest([]byte(`{"key":"abc-1", "value": "1"}`))

	checkResponseCode(t, http.StatusCreated, response.Code)

	req, _ := http.NewRequest("GET", "/searchPrefix/ab", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		if strings.Contains("body", "abc-1") == true {
			t.Errorf("Expected 1. Got %s", body)
		}
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func executePostRequest(jsonStr []byte) *httptest.ResponseRecorder {

	req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	return executeRequest(req)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
