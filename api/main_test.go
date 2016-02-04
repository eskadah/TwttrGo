package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestApi(t *testing.T){
	req, err := http.NewRequest("GET", "/search?query=mrt", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	searchResults(res, req)

	exp := `["emerita","emirate","mart","merit"]`
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
