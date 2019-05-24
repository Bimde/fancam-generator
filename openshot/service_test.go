package openshot

import (
	"os"
	"testing"
)

const (
	testUsername = "demo-cloud"
	testPassword = "demo-password"

	baseURL = "http://cloud.openshot.org/"
)

var (
	openShot *OpenShot
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	openShot = New(baseURL, testUsername, testPassword)
}

func shutdown() {
	openShot = nil
}
