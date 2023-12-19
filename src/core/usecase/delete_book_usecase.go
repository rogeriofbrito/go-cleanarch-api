package usecase

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/external"

type DeleteBookUseCase struct {
	Br external.IBookRepository
}

func (db DeleteBookUseCase) Execute(id int) error {
	return db.Br.Delete(id)
}
