package rest

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	// HTTP transport that disables checking of the TLS certificate
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
)

// QParms are the optional query parameters to include on a HTTP request
type QParms map[string]interface{}

// Headers are the optional headers to include on an HTTP request
type Headers map[string]string

// Client is a REST client that may be used to send request to a server
type Client interface {
	// Headers retrieves the optional headers to include on the REST request
	Headers() Headers
	// QParms retrieves the optional query parameters to include on the REST request
	QParms() QParms
	// Get sends a GET request to the HTTP server
	Get(url string) ([]byte, error)
}

// NewClient creates a new REST client
func NewClient(headers Headers, qparms QParms) Client {
	c := &client{headers: headers, qparms: qparms}
	return c
}

// Client is the HTTP client used to send the request to a server.
type client struct {
	headers Headers
	qparms  QParms
}

// Headers retrieves the optional headers to include on the REST request
func (c *client) Headers() Headers {
	return c.headers
}

// QParms retrieves the optional query parameters to include on the REST request
func (c *client) QParms() QParms {
	return c.qparms
}

// Get sends a GET request to the HTTP server
func (c *client) Get(url string) ([]byte, error) {
	const M = "rest.Client.Get"
	log.Debug(M, " -->")
	defer log.Debug(M, " <--")

	// Add any query paramegters to the URL
	var sb strings.Builder
	sb.Grow(100)
	sb.WriteString(url)
	if c.QParms() != nil {
		first := true
		for k, v := range c.QParms() {
			if first {
				sb.WriteString("?")
				first = false
			} else {
				sb.WriteString("&")
			}
			sb.WriteString(fmt.Sprintf("%s=%v", k, escapeString(v)))
		}
	}
	urlWithQparms := sb.String()
	log.Trace("url=" + urlWithQparms)

	// Get the http request
	log.Debug("GET url=", url)
	req, err := http.NewRequest("GET", urlWithQparms, nil)
	if err != nil {
		log.Error("failed to get the http request")
		return nil, err
	}

	// Add any custom headers
	if c.Headers() != nil && len(c.Headers()) > 0 {
		for k, v := range c.Headers() {
			req.Header.Set(k, v)
		}
	}

	// Send the request to Clash of Clans and get the response
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("failed to send the request to CoC")
		return nil, err
	}
	defer resp.Body.Close()

	// If an error status code was returned by the server, pass the error back to the invoker
	if resp.StatusCode != 200 {
		log.Error("failed to send the request to CoC, statusCode=", resp.StatusCode, ", status=", resp.Status)
		err := ErrHttp{StatusCode: resp.StatusCode, Status: resp.Status}
		return nil, err
	}

	// Read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("failed to read the body")
		return nil, err
	}
	log.Trace("response body=" + string(body))

	// All good, so return the response
	return body, nil
}

// escapeString will escape a string, otherwise it returns the value unchanged
func escapeString(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		str := value.(string)
		return url.QueryEscape(str)
	default:
		return v
	}
}
