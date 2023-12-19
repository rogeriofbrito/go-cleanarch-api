package factory

import (
	"os"

	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
)

func NewEnv() env.IEnv {
	envType := os.Getenv("ENV_TYPE") // HOST or STATIC
	switch envType {
	case "HOST":
		return env.HostEnv{}
	case "STATIC":
		return env.StaticEnv{}
	}
	return env.StaticEnv{} // default
}
