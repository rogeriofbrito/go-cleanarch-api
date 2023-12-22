package controller

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/usecase"
)

type BookModel struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

type Controller struct {
	Cb usecase.CreateBookUseCase
	Gb usecase.GetBookUseCase
	Ub usecase.UpdateBookUseCase
	Db usecase.DeleteBookUseCase
}

func (bc Controller) CreateBook(bm BookModel) (BookModel, error) {
	bd := domain.BookDomain{
		Id:    bm.Id,
		Title: bm.Title,
		Pages: bm.Pages,
	}

	bd, err := bc.Cb.Execute(bd)
	if err != nil {
		return BookModel{}, err
	}

	return BookModel{
		Id:    bd.Id,
		Title: bd.Title,
		Pages: bd.Pages,
	}, nil
}

func (bc Controller) GetBook(id int64) (BookModel, error) {
	bd, err := bc.Gb.Execute(id)
	if err != nil {
		return BookModel{}, err
	}

	return BookModel{
		Id:    bd.Id,
		Title: bd.Title,
		Pages: bd.Pages,
	}, nil
}

func (bc Controller) UpdateBook(bm BookModel) (BookModel, error) {
	bd := domain.BookDomain{
		Id:    bm.Id,
		Title: bm.Title,
		Pages: bm.Pages,
	}

	bd, err := bc.Ub.Execute(bd)
	if err != nil {
		return BookModel{}, err
	}

	return BookModel{
		Id:    bd.Id,
		Title: bd.Title,
		Pages: bd.Pages,
	}, nil
}

func (bc Controller) DeleteBook(id int64) error {
	if err := bc.Db.Execute(id); err != nil {
		return err
	}

	return nil
}
