package factory

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
	"github.com/rogeriofbrito/go-cleanarch-api/src/infra/controller"
)

func NewController() controller.Controller {
	return controller.Controller{
		Cb: NewCreateBookUseCase(),
		Gb: NewGetBookUseCase(),
		Ub: NewUpdateBookUseCase(),
	}
}

func NewControllerAdapter() controller.IControllerAdapter {
	envs := NewEnv()

	switch envs.GetHttpFramework() {
	case env.Gin:
		return NewGinControllerAdapter()
	case env.Fiber:
		return NewFiberControllerAdapter()
	}

	return NewFiberControllerAdapter() // default
}

func NewGinControllerAdapter() controller.GinControllerAdapter {
	return controller.GinControllerAdapter{
		Controller: NewController(),
		Gin:        NewGinEngine(),
	}
}

func NewFiberControllerAdapter() controller.FiberControllerAdapter {
	return controller.FiberControllerAdapter{
		Controller: NewController(),
		App:        NewFiberApp(),
	}
}
