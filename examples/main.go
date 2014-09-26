package main

import (
	"fmt"
	"net/http"

	"github.com/bradleyg/go-address"
)

func main() {
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
