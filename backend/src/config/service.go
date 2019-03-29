package config

var config map[string]interface{}

const (
	USERNAME = "username"
	PASSWORD = "password"
)

func init() {
	config = map[string]interface{}{}
	config[USERNAME] = "cloud-openshot" // TODO get from environment
	config[PASSWORD] = "cloud-password"
}

// Get returns a property from the configuration with the given key
func Get(property string) interface{} {
	return config[property]
}

// GetString returns a property from the configuration type-asserted to a string
func GetString(property string) string {
	return config[property].(string)
}

// Set allows overriding of the environment configuration programatically
func Set(key string, value interface{}) {
	config[key] = value
}
