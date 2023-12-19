package usecase

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/external"
)

type GetBookUseCase struct {
	Br external.IBookRepository
}

func (gb GetBookUseCase) Execute(id int64) (domain.BookDomain, error) {
	return gb.Br.GetById(id)
}
