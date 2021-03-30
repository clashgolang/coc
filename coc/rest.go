package coc

import (
	"github.com/clashgolang/coc/pkg/rest"
)

var (
	defaultHeaders = rest.Headers{
		"Accept": "application/json",
	}
	token string
)

// SetToken sets the token to be used on requests sent to Clash of Clans
func SetToken(t string) {
	token = t
}

// get retrieves the requested URL and return the results as a byte array.
func get(url string, qparms rest.QParms) ([]byte, error) {
	headers := rest.Headers{"Authorization": "Bearer " + token}
	for k, v := range defaultHeaders {
		headers[k] = v
	}
	client := rest.NewClient(headers, qparms)

	body, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}
