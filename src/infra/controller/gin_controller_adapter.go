package controller

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinControllerAdapter struct {
	Controller Controller
	Gin        *gin.Engine
}

// adapts to CreateBookUseCase controller function
func (gc GinControllerAdapter) CreateBookUseCase(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	request := Request{
		pathVariables: nil, // TODO: parse path variables
		params:        nil, // TODO: parse params
		headers:       nil, // TODO: parse headers
		body:          body,
	}

	b, err := gc.Controller.CreateBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

// adapts to GetBook controller function
func (gc GinControllerAdapter) GetBook(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	request := Request{
		pathVariables: map[string]string{
			"id": c.Param("id"),
		},
		params:  nil, // TODO: parse params
		headers: nil, // TODO: parse headers
		body:    body,
	}

	b, err := gc.Controller.GetBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) Start() error {
	gc.Gin.GET("/book/:id", func(c *gin.Context) {
		gc.GetBook(c)
	})

	gc.Gin.POST("/book", func(c *gin.Context) {
		gc.CreateBookUseCase(c)
	})

	return gc.Gin.Run(":3000")
}
