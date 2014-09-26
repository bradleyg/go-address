# goaddress
--
    import "github.com/bradleyg/go-address"

Go-address gets the IP address when given a *http.Request object. By passing nil
as the header argument you are asking to read the IP from r.RemoteAddr.

    addr, err := goaddress.Get(r, nil)

You can optionally pass a string to specify to look at a header rather than the
remote address. This is useful when serving requests behind a proxy. For example
Heroku passes through the remote IP in the header "X-Forwarded-For".

    addr, err := goaddress.Get(r, "X-Forwarded-For")

## Usage

#### func  Get

```go
func Get(r *http.Request, header interface{}) (string, error)
```
Get takes a request object and returns a string containing the IP address.
Header can be either nil or a string containing which header to read from.
