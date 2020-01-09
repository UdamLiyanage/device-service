package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e103f782aa554ae4e6abb8b is always available as a test document
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/devices/5e103f782aa554ae4e6abb8b", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
}

func TestInvalidDeviceGet(t *testing.T) {
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/devices/5e103f782aa554ae4e6abb8", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
