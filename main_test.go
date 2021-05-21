package main

import (
	"io/ioutil"
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

	// Test for status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %v, Got: %v", http.StatusOK, status)
	}

	// Test for Body response
	expected := "Hello World"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Unexpected Body: Expected: %v, Got: %v", actual, expected)
	}

}

func TestRouter(t *testing.T) {
	r := routesBuilder()
	mockServer := httptest.NewServer(r)

	res, err := http.Get(mockServer.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	// Test for status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %v, Got: %v", http.StatusOK, res.StatusCode)
	}

	// Test for body response
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	resString := string(b)
	expected := "Hello World"
	if resString != expected {
		t.Errorf("Unexpected Body. Expected: %v, Got: %v", expected, resString)
	}

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := routesBuilder()
	mockServer := httptest.NewServer(r)

	res, err := http.Post(mockServer.URL+"/", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test for status code
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %v", res.StatusCode)
	}

	// Test for body response
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	resString := string(b)
	expected := ""
	if resString != expected {
		t.Errorf("Unexpected Body. Expected: %v, Got: %v", expected, resString)
	}

}

func TestStaticFiles(t *testing.T) {
	r := routesBuilder()
	mockServer := httptest.NewServer(r)

	res, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// Test for status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %v, Got: %v", http.StatusOK, res.StatusCode)
	}

	// Test for header type
	contentType := res.Header.Get("Content-Type")
	expected := "text/html; charset=utf-8"
	if contentType != expected {
		t.Errorf("Unexpected Body. Expected: %v, Got: %v", expected, contentType)
	}
}
