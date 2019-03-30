package openshot

import (
	"config"
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
	config.Set(config.USERNAME, testUsername)
	config.Set(config.PASSWORD, testPassword)
	openShot = New()
}

func shutdown() {
	openShot = nil
}
