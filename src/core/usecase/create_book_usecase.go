package usecase

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/external"
)

type CreateBookUseCase struct {
	Br external.IBookRepository
}

func (cb CreateBookUseCase) Execute(book domain.BookDomain) (domain.BookDomain, error) {
	id, err := cb.Br.Save(book)
	if err != nil {
		return domain.BookDomain{}, err
	}

	book, err = cb.Br.GetById(id)
	if err != nil {
		return domain.BookDomain{}, err
	}

	return book, nil
}
