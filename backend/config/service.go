package config

var config map[string]interface{}

func init() {
	config["username"] = "cloud_openshot"
	config["password"] = "cloud_password"
}

// Get returns a property from the configuration with the given key
func Get(property string) interface{} {
	return config[property]
}
