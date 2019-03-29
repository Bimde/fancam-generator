package httputils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"config"

	"github.com/sirupsen/logrus"
)

var client http.Client

func init() {
	client = http.Client{}
}

func Get(log *logrus.Entry, path string, body *[]byte, output *interface{}) error {
	req, err := http.NewRequest("GET", path, bytes.NewBuffer(*body))
	if err != nil {
		log.Panic("error creating request ", err)
	}
	addAuth(req)
	resp, err := client.Do(req)

	if err != nil {
		log.Error("error executing request ", err)
		return err
	}

	log.Info("Response: ", resp)
	log.Info("Response Body: ", resp.Body)

	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, output)

	if err != nil {
		log.Error("error unmarshalling projects ", err)
		return err
	}

	return nil
}

func addAuth(Request *http.Request) {
	username := string(config.Get("username"))
	Request.SetBasicAuth(o.username, o.password)
}
