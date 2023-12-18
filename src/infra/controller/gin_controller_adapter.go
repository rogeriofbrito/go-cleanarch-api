package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinControllerAdapter struct {
	Controller Controller
	Gin        *gin.Engine
}

// adapts to CreateBookUseCase controller function
func (gc GinControllerAdapter) CreateBookUseCase(c *gin.Context) {
	book := new(Book)             // TODO: parse body
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	*book = gc.Controller.CreateBookUseCase(params, headers, *book)
	c.JSON(http.StatusOK, book)
}

// adapts to GetBook controller function
func (gc GinControllerAdapter) GetBook(c *gin.Context) error {
	var params map[string]string  // TODO: parse params
	var headers map[string]string // TODO: parse headers
	c.JSON(http.StatusOK, gc.Controller.GetBook(params, headers))
	return nil
}

func (gc GinControllerAdapter) Start() {
	gc.Gin.GET("/book", func(c *gin.Context) {
		gc.GetBook(c)
	})

	gc.Gin.POST("/book", func(c *gin.Context) {
		gc.CreateBookUseCase(c)
	})

	gc.Gin.Run(":3000")
}
