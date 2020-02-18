package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateDeviceAndDelete(t *testing.T) {
	device := Device{
		Name:   "Go Test Device",
		Serial: "Go_Test_Device_Serial",
	}
	body, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	createRecorder := testRequestStatusCode("POST", "/devices", body, http.StatusCreated, t)

	_ = json.NewDecoder(createRecorder.Body).Decode(&device)
	testRequestStatusCode("DELETE", "/devices/"+device.ID.Hex(), nil, http.StatusNotFound, t)
}
