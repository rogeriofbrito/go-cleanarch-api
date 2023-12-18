package controller

import (
	"github.com/gofiber/fiber/v2"
)

type FiberControllerAdapter struct {
	Controller Controller
	App        *fiber.App
}

// adapts to CreateBookUseCase controller function
func (fc FiberControllerAdapter) CreateBookUseCase(c *fiber.Ctx) error {
	request := Request{
		pathVariables: nil, // TODO: parse path variables
		params:        nil, // TODO: parse params
		headers:       nil, // TODO: parse headers
		body:          c.Request().Body(),
	}

	b, err := fc.Controller.CreateBookUseCase(request)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

// adapts to GetBook controller function
func (fc FiberControllerAdapter) GetBook(c *fiber.Ctx) error {
	request := Request{
		pathVariables: nil, // TODO: parse path variables
		params:        nil, // TODO: parse params
		headers:       nil, // TODO: parse headers
		body:          c.Request().Body(),
	}

	b, err := fc.Controller.GetBook(request)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) Start() error {
	fc.App.Get("/book", func(c *fiber.Ctx) error {
		return fc.GetBook(c)
	})

	fc.App.Post("/book", func(c *fiber.Ctx) error {
		return fc.CreateBookUseCase(c)
	})

	return fc.App.Listen(":3000")
}
