package usecase

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/external"
)

type CreateBookUseCase struct {
	Br external.IBookRepository
}

func (cb CreateBookUseCase) Execute(book domain.BookDomain) domain.BookDomain {
	return cb.Br.Save(book)
}
