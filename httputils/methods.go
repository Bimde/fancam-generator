package httputils

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Get performs a GET request using the net/http client.
// Logging is done to the provided logger
func (c *Client) Get(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Get")
	}
	return c.genericRequest(http.MethodGet, log, path, body, output)
}

// Delete performs a DELETE request using the net/http client.
// Logging is done to the provided logger
func (c *Client) Delete(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Delete")
	}
	return c.genericRequest(http.MethodDelete, log, path, body, output)
}

// Post performs a POST request using the net/http client.
// Logging is done to the provided logger
func (c *Client) Post(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Post")
	}
	return c.genericRequest(http.MethodPost, log, path, body, output)
}

// Put performs a PUT request using the net/http client.
// Logging is done to the provided logger
func (c *Client) Put(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Put")
	}
	return c.genericRequest(http.MethodPut, log, path, body, output)
}
