package httputils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"config"

	"github.com/sirupsen/logrus"
)

var client http.Client

func init() {
	client = http.Client{}
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
		log.Error("error unmarshalling projects ", err)
		return err
	}

	return nil
}

func addAuth(request *http.Request) {
	username := config.GetString(config.USERNAME)
	password := config.GetString(config.PASSWORD)
	request.SetBasicAuth(username, password)
}
