package controller

import (
	"encoding/json"
	"strconv"

	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/usecase"
)

type Request struct {
	pathVariables map[string]string
	params        map[string]string
	headers       map[string]string
	body          []byte
}

type BookModel struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

type Controller struct {
	Cb usecase.CreateBookUseCase
	Gb usecase.GetBookUseCase
	Ub usecase.UpdateBookUseCase
	Db usecase.DeleteBookUseCase
}

func (bc Controller) CreateBook(request Request) (BookModel, error) {
	bm := new(BookModel)
	err := json.Unmarshal(request.body, &bm)
	if err != nil {
		return BookModel{}, err
	}

	bd := domain.BookDomain{
		Id:    bm.Id,
		Title: bm.Title,
		Pages: bm.Pages,
	}

	bd, err = bc.Cb.Execute(bd)
	if err != nil {
		return BookModel{}, err
	}

	return BookModel{
		Id:    bd.Id,
		Title: bd.Title,
		Pages: bd.Pages,
	}, nil
}

func (bc Controller) GetBook(request Request) (BookModel, error) {
	id, err := strconv.Atoi(request.pathVariables["id"])
	if err != nil {
		return BookModel{}, err
	}

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

func (bc Controller) UpdateBook(request Request) (BookModel, error) {
	bm := new(BookModel)
	err := json.Unmarshal(request.body, &bm)
	if err != nil {
		return BookModel{}, err
	}

	bd := domain.BookDomain{
		Id:    bm.Id,
		Title: bm.Title,
		Pages: bm.Pages,
	}

	bd, err = bc.Ub.Execute(bd)
	if err != nil {
		return BookModel{}, err
	}

	return BookModel{
		Id:    bd.Id,
		Title: bd.Title,
		Pages: bd.Pages,
	}, nil
}

func (bc Controller) DeleteBook(request Request) error {
	id, err := strconv.Atoi(request.pathVariables["id"])
	if err != nil {
		return err
	}

	if err := bc.Db.Execute(id); err != nil {
		return err
	}

	return nil
}
