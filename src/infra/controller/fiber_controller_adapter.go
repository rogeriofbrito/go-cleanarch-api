package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FiberControllerAdapter struct {
	Controller Controller
	App        *fiber.App
}

func (fc FiberControllerAdapter) CreateBook(c *fiber.Ctx) error {
	b := BookModel{}
	if err := c.BodyParser(&b); err != nil {
		return err
	}

	b, err := fc.Controller.CreateBook(b)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) GetBook(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	b, err := fc.Controller.GetBook(id)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) UpdateBook(c *fiber.Ctx) error {
	b := BookModel{}
	if err := c.BodyParser(&b); err != nil {
		return err
	}

	b, err := fc.Controller.UpdateBook(b)
	if err != nil {
		return err
	}

	c.JSON(b)
	return nil
}

func (fc FiberControllerAdapter) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := fc.Controller.DeleteBook(id); err != nil {
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
