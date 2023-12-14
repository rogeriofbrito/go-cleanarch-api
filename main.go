package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeriofbrito/go-mvc/src/adapter"
	"github.com/rogeriofbrito/go-mvc/src/controller"
)

func main() {
	controller := controller.Controller{}
	webFramework := os.Getenv("WEB_FRAMEWORK")

	if webFramework == "FIBER" {
		fa := adapter.FiberController{
			Controller: controller,
			App:        fiber.New(),
		}

		fa.Start()
	} else if webFramework == "GIN" {
		ga := adapter.GinController{
			Controller: controller,
			Gin:        gin.Default(),
		}

		ga.Start()
	} else {
		fmt.Println("invalid web framework " + webFramework)
	}

}
