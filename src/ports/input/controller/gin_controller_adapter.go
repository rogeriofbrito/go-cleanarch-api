package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controller_model "github.com/rogeriofbrito/go-mvc/src/ports/input/controller/model"
)

type GinController struct {
	Controller Controller
	Gin        *gin.Engine
}

// adapts to CreateBook controller function
func (gc GinController) CreateBook(c *gin.Context) {
	book := new(controller_model.Book) // TODO: parse body
	var params map[string]string       // TODO: parse params
	var headers map[string]string      // TODO: parse headers

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	*book = gc.Controller.CreateBook(params, headers, *book)
	c.JSON(http.StatusOK, book)
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
