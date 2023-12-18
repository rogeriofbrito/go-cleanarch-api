package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
)

type MySqlBookRepository struct {
	Env env.IEnv
}

func (msbr MySqlBookRepository) Save(book domain.BookDomain) (domain.BookDomain, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		msbr.Env.GetMySqlBookUser(),
		msbr.Env.GetMySqlBookPassword(),
		msbr.Env.GetMySqlBookHost(),
		msbr.Env.GetMySqlBookPort(),
		msbr.Env.GetMySqlBookDatabase(),
	)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	_, err = db.Query("INSERT INTO book VALUES(?, ?, ?)", book.Id, book.Title, book.Pages)
	if err != nil {
		return domain.BookDomain{}, err
	}

	defer db.Close()
	return domain.BookDomain{
		Id:    book.Id,
		Title: book.Title,
		Pages: book.Pages,
	}, nil
}
