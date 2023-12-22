package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinControllerAdapter struct {
	Controller Controller
	Gin        *gin.Engine
}

func (gc GinControllerAdapter) CreateBook(c *gin.Context) {
	b := BookModel{}
	if err := c.BindJSON(&b); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	b, err := gc.Controller.CreateBook(b)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) GetBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	b, err := gc.Controller.GetBook(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) UpdateBook(c *gin.Context) {
	b := BookModel{}
	if err := c.BindJSON(&b); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	b, err := gc.Controller.UpdateBook(b)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, b)
}

func (gc GinControllerAdapter) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := gc.Controller.DeleteBook(id); err != nil {
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
