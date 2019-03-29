package openshot

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	baseURL     = "http://cloud.openshot.org"
	loggingName = "OpenShot"
)

type OpenShot struct {
	client   *http.Client
	username string
	password string
}

// New creates a new instance of OpenShot with a deafult http.Client.
func New(Username string, Password string) *OpenShot {
	return &OpenShot{&http.Client{}, Username, Password}
}

func (o *OpenShot) createReqWithAuth(method string, path string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		getLogger("createReqWithAuth").Panic("error creating request ", err)
	}
	o.addAuth(req)
	return req
}

func (o *OpenShot) addAuth(Request *http.Request) {
	Request.SetBasicAuth(o.username, o.password)
}

func getLogger(method string) *log.Entry {
	return log.WithFields(log.Fields{
		"method": fmt.Sprintf("%s#%s", loggingName, method),
	})
}
