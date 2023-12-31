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

func (msbr MySqlBookRepository) Save(book domain.BookDomain) (int64, error) {
	db, err := msbr.getDb()
	if err != nil {
		return 0, err
	}

	err = db.Ping()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO book(title, pages) VALUES(?, ?)", book.Title, book.Pages)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	defer db.Close()
	return id, nil
}

func (msbr MySqlBookRepository) GetById(id int64) (domain.BookDomain, error) {
	db, err := msbr.getDb()
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	row := db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	book := domain.BookDomain{}

	err = row.Scan(&book.Id, &book.Title, &book.Pages)
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

func (msbr MySqlBookRepository) Update(book domain.BookDomain) (domain.BookDomain, error) {
	db, err := msbr.getDb()
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	_, err = db.Exec("UPDATE book SET title = ?, pages = ? WHERE id = ?", book.Title, book.Pages, book.Id)
	if err != nil {
		return domain.BookDomain{}, err
	}

	defer db.Close()
	return book, nil
}

func (msbr MySqlBookRepository) Delete(id int64) error {
	db, err := msbr.getDb()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

func (msbr MySqlBookRepository) getDb() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		msbr.Env.GetMySqlBookUser(),
		msbr.Env.GetMySqlBookPassword(),
		msbr.Env.GetMySqlBookHost(),
		msbr.Env.GetMySqlBookPort(),
		msbr.Env.GetMySqlBookDatabase(),
	)

	return sql.Open("mysql", url)
}
