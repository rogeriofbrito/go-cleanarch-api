package factory

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/usecase"

func NewCreateBookUseCase() usecase.CreateBookUseCase {
	return usecase.CreateBookUseCase{
		Br: NewBookRepository(),
	}
}

func NewGetBookUseCase() usecase.GetBookUseCase {
	return usecase.GetBookUseCase{
		Br: NewBookRepository(),
	}
}
