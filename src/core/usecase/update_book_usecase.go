package usecase

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/external"
)

type UpdateBookUseCase struct {
	Br external.IBookRepository
}

func (ub UpdateBookUseCase) Execute(book domain.BookDomain) (domain.BookDomain, error) {
	return ub.Br.Update(book)
}
