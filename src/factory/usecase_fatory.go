package factory

import "github.com/rogeriofbrito/go-mvc/src/core/usecase"

func NewCreateBook() usecase.CreateBook {
	return usecase.CreateBook{
		Br: NewBookRepository(),
	}
}
