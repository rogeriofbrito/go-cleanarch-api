package env

import (
	"os"
	"strconv"
)

type HostEnv struct{}

func (he HostEnv) GetHttpFramework() HttpFrameworkType {
	return HttpFrameworkType(os.Getenv("HTTP_FRAMEWORK"))
}

func (he HostEnv) GetDatabaseType() DatabaseType {
	return DatabaseType(os.Getenv("DATABESE"))
}

func (he HostEnv) GetMySqlBookDatabase() string {
	return os.Getenv("MYSQL_BOOK_DATABASE")
}

func (he HostEnv) GetMySqlBookUser() string {
	return os.Getenv("MYSQL_BOOK_USER")
}

func (he HostEnv) GetMySqlBookPassword() string {
	return os.Getenv("MYSQL_BOOK_PASSWORD")
}

func (he HostEnv) GetMySqlBookHost() string {
	return os.Getenv("MYSQL_BOOK_HOST")
}

func (he HostEnv) GetMySqlBookPort() int {
	port, err := strconv.Atoi(os.Getenv("MYSQL_BOOK_PORT"))
	if err != nil {
		panic(err)
	}
	return port
}

func (he HostEnv) GetPostgresBookDatabase() string {
	return os.Getenv("POSTGRES_BOOK_DATABASE")
}

func (he HostEnv) GetPostgresBookUser() string {
	return os.Getenv("POSTGRES_BOOK_USER")
}

func (he HostEnv) GetPostgresBookPassword() string {
	return os.Getenv("POSTGRES_BOOK_PASSWORD")
}

func (he HostEnv) GetPostgresBookHost() string {
	return os.Getenv("POSTGRES_BOOK_HOST")
}

func (he HostEnv) GetPostgresBookPort() int {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_BOOK_PORT"))
	if err != nil {
		panic(err)
	}
	return port
}
