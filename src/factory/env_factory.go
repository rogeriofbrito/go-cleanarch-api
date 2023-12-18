package factory

import "github.com/rogeriofbrito/go-mvc/src/env"

func NewEnv() env.IEnv {
	return env.StaticEnv{} // default
}
