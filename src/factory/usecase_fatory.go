package factory

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/usecase"

func NewCreateBook() usecase.CreateBook {
	return usecase.CreateBook{
		Br: NewBookRepository(),
	}
}
