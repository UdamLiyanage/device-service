package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	req, _ := http.NewRequest("PUT", "/devices/000000000000000000000000", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}
	var response map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &response)
	value, exists := response["MatchedCount"]
	if !exists {
		t.Errorf("Wrong Response Format")
	}
	count, _ := strconv.Atoi(value)
	if count != 0 {
		t.Errorf("Operation Not Working Properly!")
	}
}
