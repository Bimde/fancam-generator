package openshot

import (
	"os"
	"testing"
)

const (
	testUsername = "demo-cloud"
	testPassword = "demo-password"
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
	openShot = New(testUsername, testPassword)
}

func shutdown() {
	openShot = nil
}
