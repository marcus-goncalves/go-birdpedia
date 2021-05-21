package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHome(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handleHome)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %v, Got: %v", http.StatusOK, status)
	}

	expected := "Hello World"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Unexpected Body: Expected: %v, Got: %v", actual, expected)
	}

}
