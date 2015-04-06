package goaddress

import (
	"log"
	"net/http"
	"testing"
)

var req *http.Request

func init() {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = r
}

func TestGetAddrWithPort(t *testing.T) {
	req.RemoteAddr = "0.0.0.0:80"

	address, err := Get(req, nil)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "0.0.0.0" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "0.0.0.0", address)
	}
}

func TestGetAddrWithoutPort(t *testing.T) {
	req.RemoteAddr = "0.0.0.0"

	address, err := Get(req, nil)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "0.0.0.0" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "0.0.0.0", address)
	}
}

func TestGetAddrWithHeader(t *testing.T) {
	req.Header.Set("HTTP_X_FORWARDED_FOR", "1.1.1.1:80, 2.2.2.2:80")

	address, err := Get(req, "HTTP_X_FORWARDED_FOR")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "2.2.2.2" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "2.2.2.2", address)
	}
}

func TestGetAddrWithNoAddressShouldFallback(t *testing.T) {
	address, err := Get(req, "MISSING_ADDRESS")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "0.0.0.0" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "0.0.0.0", address)
	}
}
