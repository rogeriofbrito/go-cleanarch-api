package factory

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	return fiber.New()
}
