package external

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"

type IBookRepository interface {
	Save(book domain.Book) domain.Book
}
