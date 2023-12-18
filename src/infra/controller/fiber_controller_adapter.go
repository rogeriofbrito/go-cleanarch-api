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
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers
	c.JSON(fc.Controller.CreateBookUseCase(params, headers, *book))
	return nil
}

// adapts to GetBook controller function
func (fc FiberControllerAdapter) GetBook(c *fiber.Ctx) error {
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers
	c.JSON(fc.Controller.GetBook(params, headers))
	return nil
}

func (fc FiberControllerAdapter) Start() {
	fc.App.Get("/book", func(c *fiber.Ctx) error {
		return fc.GetBook(c)
	})

	fc.App.Post("/book", func(c *fiber.Ctx) error {
		return fc.CreateBookUseCase(c)
	})

	fc.App.Listen(":3000")
}
