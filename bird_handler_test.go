package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBirdHandler(t *testing.T) {
	birds = []Bird{
		{"sparrow", "Maldito passaro pequeno"},
	}

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal()
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(getBirdHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status Code. Got %v, Want %v", status, http.StatusOK)
	}

	expected := Bird{"sparrow", "Maldito passaro pequeno"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]
	if actual != expected {
		t.Errorf("Unexpected Body. Got %v, expected %v", actual, expected)
	}
}

func createBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "careca da poha")

	return &form
}

func TestCreateBirdHandler(t *testing.T) {
	birds = []Bird{
		{"eagle", "careca da poha"},
	}

	form := createBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(createBirdHandler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("Wrong status code. Got %v, want %v", status, http.StatusOK)
	}

	expected := Bird{"eagle", "careca da poha"}
	actual := birds[1]
	if actual != expected {
		t.Errorf("Unexpected Body. Got %v, Want %v", actual, expected)
	}

}
