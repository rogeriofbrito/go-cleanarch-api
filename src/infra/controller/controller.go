package controller

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/usecase"
)

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

type Controller struct {
	Cb usecase.CreateBook
}

func (bc Controller) CreateBook(params map[string]string, headers map[string]string, body Book) Book {
	book := domain.Book{
		Id:    body.Id,
		Title: body.Title,
		Pages: body.Pages,
	}
	book = bc.Cb.Execute(book)

	return Book{
		Id:    book.Id,
		Title: book.Title,
		Pages: book.Pages,
	}
}

func (bc Controller) GetBook(params map[string]string, headers map[string]string) Book {
	return Book{}
}
