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

func TestgetAddrWithoutPort(t *testing.T) {
	req.RemoteAddr = "0.0.0.0"

	address, err := Get(req, nil)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "0.0.0.0" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "0.0.0.0", address)
	}
}

func TestgetAddrWithHeader(t *testing.T) {
	req.Header.Set("HTTP_X_FORWARDED_FOR", "1.1.1.1:80, 2.2.2.2:80")

	address, err := Get(req, "HTTP_X_FORWARDED_FOR")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	if address != "1.1.1.1" {
		t.Fatalf("Address doesn't match. Expected %s, Actual %s", "1.1.1.1", address)
	}
}

func TestgetAddrWithNoAddress(t *testing.T) {
	_, err := Get(req, "MISSING_ADDRESS")
	if err == nil {
		t.Fatalf("A missing address should return an error")
	}
}

// Examples

func ExampleGet() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		address, err := goaddress.Get(r, nil)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Your IP: %s", address)
	})

	http.ListenAndServe(":8080", mux)
}
