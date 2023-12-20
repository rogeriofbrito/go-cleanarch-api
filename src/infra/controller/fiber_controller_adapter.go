package controller

import (
	"github.com/gofiber/fiber/v2"
)

type FiberControllerAdapter struct {
	Controller Controller
	App        *fiber.App
}

func (fc FiberControllerAdapter) CreateBook(c *fiber.Ctx) error {
	request := Request{
		body: c.Request().Body(),
	}

	b, err := fc.Controller.CreateBook(request)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) GetBook(c *fiber.Ctx) error {
	request := Request{
		pathVariables: map[string]string{
			"id": c.Params("id"),
		},
		body: c.Request().Body(),
	}

	b, err := fc.Controller.GetBook(request)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) UpdateBook(c *fiber.Ctx) error {
	request := Request{
		body: c.Request().Body(),
	}

	b, err := fc.Controller.UpdateBook(request)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) DeleteBook(c *fiber.Ctx) error {
	request := Request{
		pathVariables: map[string]string{
			"id": c.Params("id"),
		},
		body: nil,
	}

	err := fc.Controller.DeleteBook(request)
	if err != nil {
		return err
	}

	return nil
}

func (fc FiberControllerAdapter) Start() error {
	fc.App.Get("/book/:id", func(c *fiber.Ctx) error {
		return fc.GetBook(c)
	})

	fc.App.Post("/book", func(c *fiber.Ctx) error {
		return fc.CreateBook(c)
	})

	fc.App.Put("/book", func(c *fiber.Ctx) error {
		return fc.UpdateBook(c)
	})

	fc.App.Delete("/book/:id", func(c *fiber.Ctx) error {
		return fc.DeleteBook(c)
	})

	return fc.App.Listen("127.0.0.1:3000")
}
