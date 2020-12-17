package env

import (
	r "template.clevecord.me/server/repositories/env"
)

// LookupEnv tries to look up a value, putting the value inside the second parameter.
// If an error was found, the error will not be nil.
func LookupEnv(key string) (string, error) {
	return r.LookupEnv(key)
}
