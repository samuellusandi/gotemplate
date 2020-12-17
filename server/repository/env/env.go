package env

import (
	"errors"

	"template.clevecord.me/server/constants"
)

// LookupEnv tries to look up an environment variable as defined in
// .env file. It will return an error when the variable is not found.
func LookupEnv(key string) (string, error) {
	value := constants.Lookup(key)
	if value == constants.UnknownValue {
		return "", errors.New("environment variable not found")
	}
	return value, nil
}
