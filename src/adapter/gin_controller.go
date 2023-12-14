package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofbrito/go-mvc/src/controller"
	"github.com/rogeriofbrito/go-mvc/src/model"
)

type GinController struct {
	Controller controller.IController
	Gin        *gin.Engine
}

// adapts to CreateBook controller function
func (gc GinController) CreateBook(c *gin.Context) {
	book := new(model.Book)       // TODO: parse body
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers
	c.JSON(http.StatusOK, gc.Controller.CreateBook(params, headers, *book))
}

// adapts to GetBook controller function
func (gc GinController) GetBook(c *gin.Context) error {
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers
	c.JSON(http.StatusOK, gc.Controller.GetBook(params, headers))
	return nil
}

func (gc GinController) Start() {
	gc.Gin.GET("/book", func(c *gin.Context) {
		gc.GetBook(c)
	})

	gc.Gin.POST("/book", func(c *gin.Context) {
		gc.CreateBook(c)
	})

	gc.Gin.Run(":3000")
}
