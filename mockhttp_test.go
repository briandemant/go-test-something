//go:build unit
// +build unit

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "status:ok")
	}
	req := httptest.NewRequest("GET", "https://briandemant.dk/status", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if 200 != resp.StatusCode {
		t.Fatal("why did you fail!")
	}
	if string(body) != "status:ok" {
		t.Fatalf("expected 'status:ok' got %v", string(body))
	}
}
