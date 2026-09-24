package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handleInformation(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("expected ststus 200, got %d", rr.Code)

	}

}
