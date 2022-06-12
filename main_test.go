package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_get(t *testing.T) {
	request, error := http.NewRequest("GET", "/api.uktradeinfo.com/Commodity(0)?$select=Cn8Code", nil)
	if error != nil {
		t.Fatal(error)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("status: expected %v actual %v", http.StatusOK, status)
	}

	expected := `{"@odata.context":"https://api.uktradeinfo.com/$metadata#Commodity(Cn8Code)/$entity","Cn8Code":"00"}`
	actual := strings.Trim(recorder.Body.String(), "\n")

	if expected != actual {
		t.Errorf("body: expected %v actual %v", expected, actual)
	}
}
