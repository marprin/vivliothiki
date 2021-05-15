package vivliothiki

import (
	"os"
)

// Environment List
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var env string

func init() {
	env = os.Getenv("STAGE")
	if env == "" {
		env = EnvDevelopment
	}
}

// GetEnv return the string of current environment flag
func GetEnv() string {
	return env
}
