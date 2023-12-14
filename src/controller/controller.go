package controller

import "github.com/rogeriofbrito/go-mvc/src/model"

type IController interface {
	CreateBook(params map[string]string, headers map[string]string, body model.Book) model.Book
	GetBook(params map[string]string, headers map[string]string) model.Book
}

type Controller struct {
}

func (bc Controller) CreateBook(params map[string]string, headers map[string]string, body model.Book) model.Book {
	return model.Book{}
}

func (bc Controller) GetBook(params map[string]string, headers map[string]string) model.Book {
	return model.Book{}
}
