package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e103f782aa554ae4e6abb8b is always available as a test document
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/devices/5e103f782aa554ae4e6abb8b")

	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}
}

func TestInvalidDeviceGet(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/devices/5e103f782aa554ae4e6abb8")

	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", resp.StatusCode)
	}
}
