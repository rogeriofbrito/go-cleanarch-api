package factory

import "github.com/rogeriofbrito/go-cleanarch-api/src/env"

func NewEnv() env.IEnv {
	return env.StaticEnv{} // default
}
