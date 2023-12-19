package usecase

import "github.com/rogeriofbrito/go-cleanarch-api/src/core/external"

type DeleteBookUseCase struct {
	Br external.IBookRepository
}

func (db DeleteBookUseCase) Execute(id int64) error {
	_, err := db.Br.GetById(id)
	if err != nil {
		return err
	}

	return db.Br.Delete(id)
}
