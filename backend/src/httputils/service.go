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

// Get Performs a GET request using the net/http client.
// Logging is done to the provided logger
func Get(log *logrus.Entry, path string, body io.Reader, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Get")
	}
	return genericRequest(http.MethodGet, log, path, body, output)
}

// Delete Performs a DELETE request using the net/http client.
// Logging is done to the provided logger
func Delete(log *logrus.Entry, path string, body io.Reader, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Delete")
	}
	return genericRequest(http.MethodDelete, log, path, body, output)
}

func genericRequest(method string, log *logrus.Entry, path string, body io.Reader, output interface{}) error {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Panic("error creating request ", err)
	}
	addAuth(req)
	res, err := client.Do(req)

	if err != nil {
		log.Error("error executing request ", err)
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Error("error request response status code ", res.StatusCode)
		return fmt.Errorf("error executing request to %s, status: %d", path, res.StatusCode)
	}

	log.Debug("Response: ", res)
	log.Debug("Response Body: ", res.Body)

	if output == nil {
		// Output doesn't need to be unmarshalled. This is needed for dealing with http code
		// 204 No Content responses
		return nil
	}

	bytes, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bytes, output)

	if err != nil {
		log.Error("error unmarshalling response ", err)
		return err
	}

	return nil
}

// Post Performs a POST request using the net/http client.
// Logging is done to the provided logger
func Post(log *logrus.Entry, path string, body interface{}, output interface{}) error {
	if log == nil {
		log = logrus.WithField("method", "httputils#Get")
	}

	data, err := json.Marshal(body)
	if err != nil {
		log.Error("error marshalling project ", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(data))
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

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("error reading response body ", err)
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Error("error request response status", res)
		log.Error("response body ", string(resBody))
		return fmt.Errorf("error executing request to %s, status: %d", path, res.StatusCode)
	}

	err = json.Unmarshal(resBody, output)
	if err != nil {
		log.WithField("responseBody", string(resBody)).Error("error unmarshalling response ", err)
		return err
	}

	return nil
}

func addAuth(request *http.Request) {
	username := config.GetString(config.USERNAME)
	password := config.GetString(config.PASSWORD)
	request.SetBasicAuth(username, password)
}
