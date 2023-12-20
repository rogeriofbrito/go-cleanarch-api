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

func (gc GinControllerAdapter) CreateBook(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	request := Request{
		body: body,
	}

	b, err := gc.Controller.CreateBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) GetBook(c *gin.Context) {
	request := Request{
		pathVariables: map[string]string{
			"id": c.Param("id"),
		},
		body: nil,
	}

	b, err := gc.Controller.GetBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) UpdateBook(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	request := Request{
		body: body,
	}

	b, err := gc.Controller.UpdateBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) DeleteBook(c *gin.Context) {
	request := Request{
		pathVariables: map[string]string{
			"id": c.Param("id"),
		},
		body: nil,
	}

	err := gc.Controller.DeleteBook(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (gc GinControllerAdapter) Start() error {
	gc.Gin.GET("/book/:id", func(c *gin.Context) {
		gc.GetBook(c)
	})

	gc.Gin.POST("/book", func(c *gin.Context) {
		gc.CreateBook(c)
	})

	gc.Gin.PUT("/book", func(c *gin.Context) {
		gc.UpdateBook(c)
	})

	gc.Gin.DELETE("/book/:id", func(c *gin.Context) {
		gc.DeleteBook(c)
	})

	return gc.Gin.Run("127.0.0.1:3000")
}
