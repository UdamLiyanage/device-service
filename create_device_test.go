package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDeviceAndDelete(t *testing.T) {
	device := Device{
		Name:           "Go Test Device",
		Serial:         "Go_Test_Device_Serial",
		Configurations: nil,
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	r := newRouter()
	createRecorder := httptest.NewRecorder()
	createRequest, _ := http.NewRequest("POST", "/devices", bytes.NewBuffer(body))
	r.ServeHTTP(createRecorder, createRequest)
	if createRecorder.Code != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", createRecorder.Code)
	}

	_ = json.NewDecoder(createRecorder.Body).Decode(&device)
	deleteRequest, _ := http.NewRequest("DELETE", "/devices/"+device.ID.Hex(), nil)
	deleteRecorder := httptest.NewRecorder()
	r.ServeHTTP(deleteRecorder, deleteRequest)
	if deleteRecorder.Code != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", deleteRecorder.Code)
	}
}
