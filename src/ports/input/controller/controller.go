package controller

import (
	"github.com/rogeriofbrito/go-mvc/src/core/domain"
	"github.com/rogeriofbrito/go-mvc/src/core/usecase"
	controller_model "github.com/rogeriofbrito/go-mvc/src/ports/input/controller/model"
)

type Controller struct {
	Cb usecase.CreateBook
}

func (bc Controller) CreateBook(params map[string]string, headers map[string]string, body controller_model.Book) controller_model.Book {
	book := domain.Book{
		Id:    body.Id,
		Title: body.Title,
		Pages: body.Pages,
	}
	book = bc.Cb.Execute(book)

	return controller_model.Book{
		Id:    book.Id,
		Title: book.Title,
		Pages: book.Pages,
	}
}

func (bc Controller) GetBook(params map[string]string, headers map[string]string) controller_model.Book {
	return controller_model.Book{}
}
