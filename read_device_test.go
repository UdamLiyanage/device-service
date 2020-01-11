package main

import (
	"net/http"
	"testing"
)

func TestValidDeviceGet(t *testing.T) {
	//Object ID - 5e103f782aa554ae4e6abb8b is always available as a test document
	testRequestStatusCode("GET", "/devices/5e103f782aa554ae4e6abb8b", nil, http.StatusOK, t)
}

func TestInvalidDeviceGet(t *testing.T) {
	testRequestStatusCode("GET", "/devices/000000000000000000000000", nil, http.StatusNotFound, t)
}
