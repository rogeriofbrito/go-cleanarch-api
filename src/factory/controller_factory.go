package factory

import (
	"github.com/rogeriofbrito/go-mvc/src/adapter"
	"github.com/rogeriofbrito/go-mvc/src/env"
)

func NewController() adapter.Controller {
	return adapter.Controller{
		Cb: NewCreateBook(),
	}
}

func NewControllerAdapter() adapter.IControllerAdapter {
	envs := NewEnv()

	switch envs.GetHttpFramework() {
	case env.Gin:
		return NewGinController()
	case env.Fiber:
		return NewFiberController()
	}

	return NewFiberController() // default
}

func NewGinController() adapter.GinController {
	return adapter.GinController{
		Controller: NewController(),
		Gin:        NewGinEngine(),
	}
}

func NewFiberController() adapter.FiberController {
	return adapter.FiberController{
		Controller: NewController(),
		App:        NewFiberApp(),
	}
}
