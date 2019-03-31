// Package httputils provides helper methods wrapping net/http.
// If this package is going to be exported for public use, the panic logs should
// be changed to errors and the dependency to config should be removed.
package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"config"

	"github.com/sirupsen/logrus"
)

const requestTimeoutSeconds = 30

var client http.Client

func init() {
	client = http.Client{
		Timeout: time.Second * requestTimeoutSeconds,
	}
}

// Get performs a GET request using the net/http client.
// Logging is done to the provided logger
func Get(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Get")
	}
	return genericRequest(http.MethodGet, log, path, body, output)
}

// Delete performs a DELETE request using the net/http client.
// Logging is done to the provided logger
func Delete(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Delete")
	}
	return genericRequest(http.MethodDelete, log, path, body, output)
}

// Post performs a POST request using the net/http client.
// Logging is done to the provided logger
func Post(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Post")
	}
	return genericRequest(http.MethodPost, log, path, body, output)
}

// Put performs a PUT request using the net/http client.
// Logging is done to the provided logger
func Put(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Put")
	}
	return genericRequest(http.MethodPut, log, path, body, output)
}

func genericRequest(method string, log *logrus.Entry, path string, input interface{}, output interface{}) error {
	var body io.Reader
	if input == nil {
		body = nil
	} else {
		data, err := json.Marshal(input)
		if err != nil {
			log.Error("error marshalling input ", err)
			return err
		}
		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Panic("error creating request ", err)
	}
	addAuth(req)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		log.Error("error executing request ", err)
		return err
	}

	log.Debug("Response: ", res)

	var bytes []byte
	if res.ContentLength > 0 {
		bytes, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Error("error reading response body ", err)
			return err
		}
		log.Debug("Response Body: ", string(bytes))
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.WithField("statusCode", res.StatusCode).Error("error request response body ", string(bytes))
		return fmt.Errorf("error executing request to %s, status: %d", path, res.StatusCode)
	}

	if output != nil {
		err = json.Unmarshal(bytes, output)
		if err != nil {
			log.Error("error unmarshalling response ", err)
			return err
		}
	}

	return nil
}

func addAuth(request *http.Request) {
	username := config.GetString(config.Username)
	password := config.GetString(config.Password)
	request.SetBasicAuth(username, password)
}
