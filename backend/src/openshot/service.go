package openshot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	baseURL     = "http://cloud.openshot.org"
	loggingName = "openshot"
)

type OpenShot struct {
}

// New creates a new instance of OpenShot with a deafult http.Client.
func New() *OpenShot {
	return &OpenShot{}
}

func getLogger(method string) *log.Entry {
	return log.WithFields(log.Fields{
		"method": fmt.Sprintf("%s#%s", loggingName, method),
	})
}