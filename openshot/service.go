// Package openshot is a Go SDK for the OpenShot Cloud API.
package openshot

import (
	"fmt"

	"github.com/Bimde/httputils"
	log "github.com/sirupsen/logrus"
)

const (
	loggingName = "openshot"
)

// OpenShot is the main entry point into the sdk
type OpenShot struct {
	BaseURL string
	http    *httputils.Client
}

// New creates a new instance of OpenShot with default settings
func New(BaseURL string, Username string, Password string) *OpenShot {
	return &OpenShot{BaseURL: BaseURL, http: httputils.New(Username, Password)}
}

func getLogger(method string) *log.Entry {
	return log.WithFields(log.Fields{
		"method": fmt.Sprintf("%s#%s", loggingName, method),
	})
}
