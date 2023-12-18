package factory

import (
	core_external "github.com/rogeriofbrito/go-mvc/src/core/external"
	"github.com/rogeriofbrito/go-mvc/src/env"
	"github.com/rogeriofbrito/go-mvc/src/port"
)

func NewBookRepository() core_external.IBookRepository {
	envs := NewEnv()

	switch envs.GetDatabaseType() {
	case env.MySql:
		return port.MySqlBookRepository{
			Env: NewEnv(),
		}
	case env.Postgres:
		return port.PostgresBookRepository{
			Env: NewEnv(),
		}
	}

	return port.MySqlBookRepository{} // default
}
