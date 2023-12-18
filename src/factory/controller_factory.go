package factory

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
	"github.com/rogeriofbrito/go-cleanarch-api/src/infra/controller"
)

func NewController() controller.Controller {
	return controller.Controller{
		Cb: NewCreateBook(),
	}
}

func NewControllerAdapter() controller.IControllerAdapter {
	envs := NewEnv()

	switch envs.GetHttpFramework() {
	case env.Gin:
		return NewGinController()
	case env.Fiber:
		return NewFiberController()
	}

	return NewFiberController() // default
}

func NewGinController() controller.GinController {
	return controller.GinController{
		Controller: NewController(),
		Gin:        NewGinEngine(),
	}
}

func NewFiberController() controller.FiberController {
	return controller.FiberController{
		Controller: NewController(),
		App:        NewFiberApp(),
	}
}
