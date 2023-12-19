package external

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"

type IBookRepository interface {
	Save(book domain.BookDomain) (int64, error)
	GetById(id int64) (domain.BookDomain, error)
	Update(book domain.BookDomain) (domain.BookDomain, error)
	Delete(id int64) error
}
