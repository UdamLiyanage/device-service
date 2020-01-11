package main

import (
	"encoding/json"
	"net/http"
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
	testRequestStatusCode("PUT", "/devices/5e103f782aa554ae4e6abb8b", body, http.StatusOK, t)
	/*r := newRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/devices/5e103f782aa554ae4e6abb8b", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Status should be 200, got %d", w.Code)
	}*/
}

func TestUpdateDeviceInvalid(t *testing.T) {
	device := Device{
		Name: "Updated Name",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/devices/000000000000000000000000", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}
