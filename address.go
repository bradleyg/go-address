// Go-address gets the IP address when given a *http.Request object. By passing
// "nil" as the "header" argument you are asking to read the IP from r.RemoteAddr.
//
//  addr, err := GetAddr(r, nil)
//
// You can optionally pass a string to specify to look at a header rather than the remote
// address. This is useful for when serving requests behind a proxy. For example
// Heroku passes through the remote IP in the header "X-Forwarded-For".
//
//  addr, err := GetAddr(r, "X-Forwarded-For")
//
package goaddress

import (
	"errors"
	"net/http"
	"strings"
)

// GetAddr takes a request object and returns a string containing the IP address.
// Header can be either nil or a string containing which header to read from.
func Get(r *http.Request, header interface{}) (string, error) {
	var value string

	switch header.(type) {
	case string:
		value = r.Header.Get(header.(string))
	default:
		value = r.RemoteAddr
	}

	addresses := strings.Split(value, ",")
	address := strings.TrimSpace(addresses[0])

	idx := strings.LastIndex(address, ":")
	if idx != -1 {
		address = address[:idx]
	}

	if address == "" {
		err := errors.New("Could not read address")
		return "", err
	}

	return address, nil
}
