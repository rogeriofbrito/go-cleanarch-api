package external

import "github.com/rogeriofbrito/go-mvc/src/core/domain"

type IBookRepository interface {
	Save(book domain.Book) domain.Book
}
