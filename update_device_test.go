package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateDeviceValid(t *testing.T) {
	device := Device{
		Name: "Updated Name",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/devices/5e103f782aa554ae4e6abb8b", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
}

func TestUpdateDeviceInvalid(t *testing.T) {
	device := Device{
		Name: "Updated Name",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/devices/5e103f782aa554ae4e6abb8", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", w.Code)
	}
}
