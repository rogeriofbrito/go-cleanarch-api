package external

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"

type IBookRepository interface {
	Save(book domain.BookDomain) (domain.BookDomain, error)
}
