// Package httputils is a net/http wrapper for easy outbound http requests.
package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const requestTimeoutSeconds = 30

var client http.Client

func init() {
	client = http.Client{
		Timeout: time.Second * requestTimeoutSeconds,
	}
}

// Client provides storage for data that doesn't need to change between requests
type Client struct {
	Username string
	Password string
}

// New creates and returns a new Client with custom properties
func New(username string, password string) *Client {
	return &Client{Username: username, Password: password}
}

func (c *Client) genericRequest(method string, log *logrus.Entry, path string, input interface{}, output interface{}) error {
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
		log.Error("error creating request ", err)
		return err
	}
	c.addAuth(req)
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

func (c *Client) addAuth(request *http.Request) {
	if c.Username != "" && c.Password != "" {
		request.SetBasicAuth(c.Username, c.Password)
	}
}
