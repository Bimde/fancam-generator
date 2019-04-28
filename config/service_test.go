package config

import "testing"

const (
	propertyKey   = "random_test_property_key"
	propertyValue = "random_test_property_value"
)

func TestSetAndGetProperty(t *testing.T) {
	Set(propertyKey, propertyValue)
	if Get(propertyKey) != propertyValue {
		t.Error("Setting and getting property failed")
	}
}
