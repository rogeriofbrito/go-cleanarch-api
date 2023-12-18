package factory

import (
	"github.com/rogeriofbrito/go-cleanarch-api/src/core/external"
	"github.com/rogeriofbrito/go-cleanarch-api/src/env"
	"github.com/rogeriofbrito/go-cleanarch-api/src/infra/repository"
)

func NewBookRepository() external.IBookRepository {
	envs := NewEnv()

	switch envs.GetDatabaseType() {
	case env.MySql:
		return repository.MySqlBookRepository{
			Env: NewEnv(),
		}
	case env.Postgres:
		return repository.PostgresBookRepository{
			Env: NewEnv(),
		}
	}

	return repository.MySqlBookRepository{} // default
}
