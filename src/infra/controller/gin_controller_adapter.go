package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinControllerAdapter struct {
	Controller Controller
	Gin        *gin.Engine
}

// adapts to CreateBookUseCase controller function
func (gc GinControllerAdapter) CreateBookUseCase(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	request := Request{
		pathVariables: nil, // TODO: parse path variables
		params:        nil, // TODO: parse params
		headers:       nil, // TODO: parse headers
		body:          body,
	}

	b, err := gc.Controller.CreateBookUseCase(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, b)
}

// adapts to GetBook controller function
func (gc GinControllerAdapter) GetBook(c *gin.Context) error {
	body, _ := ioutil.ReadAll(c.Request.Body)
	request := Request{
		pathVariables: nil, // TODO: parse path variables
		params:        nil, // TODO: parse params
		headers:       nil, // TODO: parse headers
		body:          body,
	}

	b, err := gc.Controller.GetBook(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, b)
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
