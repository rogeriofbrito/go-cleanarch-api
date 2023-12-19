package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/domain"
	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
)

type PostgresBookRepository struct {
	Env env.IEnv
}

func (pbr PostgresBookRepository) Save(book domain.BookDomain) (domain.BookDomain, error) {
	db, err := pbr.getDb()
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	_, err = db.Exec("INSERT INTO book VALUES($1, $2, $3)", book.Id, book.Title, book.Pages)

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

func (pbr PostgresBookRepository) GetById(id int) (domain.BookDomain, error) {
	db, err := pbr.getDb()
	if err != nil {
		return domain.BookDomain{}, err
	}

	err = db.Ping()
	if err != nil {
		return domain.BookDomain{}, err
	}

	row := db.QueryRow("SELECT * FROM book WHERE id = $1", id)
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

func (pbr PostgresBookRepository) getDb() (*sql.DB, error) {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pbr.Env.GetPostgresBookHost(),
		pbr.Env.GetPostgresBookPort(),
		pbr.Env.GetPostgresBookUser(),
		pbr.Env.GetPostgresBookPassword(),
		pbr.Env.GetPostgresBookDatabase())

	return sql.Open("postgres", url)
}
