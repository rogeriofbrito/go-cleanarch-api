package usecase

import (
	"github.com/rogeriofbrito/go-mvc/src/core/domain"
	core_external "github.com/rogeriofbrito/go-mvc/src/core/external"
)

type CreateBook struct {
	Br core_external.IBookRepository
}

func (cb CreateBook) Execute(book domain.Book) domain.Book {
	return cb.Br.Save(book)
}
