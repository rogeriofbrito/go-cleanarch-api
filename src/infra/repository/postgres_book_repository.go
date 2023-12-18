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

func (pbr PostgresBookRepository) Save(book domain.Book) domain.Book {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pbr.Env.GetPostgresBookHost(),
		pbr.Env.GetPostgresBookPort(),
		pbr.Env.GetPostgresBookUser(),
		pbr.Env.GetPostgresBookPassword(),
		pbr.Env.GetPostgresBookDatabase())
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO book VALUES($1, $2, $3)", book.Id, book.Title, book.Pages)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	return domain.Book{
		Id:    book.Id,
		Title: book.Title,
		Pages: book.Pages,
	}
}
