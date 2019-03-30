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
	req, err := http.NewRequest("GET", path, body)
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
		return fmt.Errorf("error executing request to %s", path)
	}

	log.Debug("Response: ", res)
	log.Debug("Response Body: ", res.Body)

	bytes, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bytes, output)

	if err != nil {
		log.Error("error unmarshalling response ", err)
		return err
	}

	return nil
}

// Get Performs a POST request using the net/http client.
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

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(data))
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
		return fmt.Errorf("error executing request to %s", path)
	}

	err = json.Unmarshal(resBody, output)
	if err != nil {
		log.WithField("responseBody", string(resBody)).Error("error unmarshalling response ", err)
	}

	return nil
}

func addAuth(request *http.Request) {
	username := config.GetString(config.USERNAME)
	password := config.GetString(config.PASSWORD)
	request.SetBasicAuth(username, password)
}
