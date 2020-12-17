package constants

import (
	"os"
)

// DefaultPort is 3000 - read when APP_PORT is not found
const DefaultPort = "3000"

// AppDescriptionEnvKey is the Key Constant to retrieve the Application Description from Env.
const AppDescriptionEnvKey = "APP_DESCRIPTION"

// AppEnvironmentEnvKey is the Key Constant to check if the application should be run on development or production (or others)
const AppEnvironmentEnvKey = "APP_ENV"

// AppPortEnvKey is the Key Constant for which port the Application should be run on.
const AppPortEnvKey = "APP_PORT"

// AppVersionEnvKey is the Key Constant to check which version the application is running on.
const AppVersionEnvKey = "APP_VERSION"

// UnknownValue is the result we will return when the value is unknown
const UnknownValue = "unknown"

// Lookup tries to look up a value; returns UnknownValue when value is not found.
func Lookup(key string, custom ...string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	if len(custom) > 0 {
		return custom[0]
	}
	return UnknownValue
}
