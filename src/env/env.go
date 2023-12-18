package env

type HttpFrameworkType string
type DatabaseType string

const (
	Fiber    HttpFrameworkType = "FIBER"
	Gin      HttpFrameworkType = "GIN"
	MySql    DatabaseType      = "MYSQL"
	Postgres DatabaseType      = "POSTGRES"
)

type IEnv interface {
	GetHttpFramework() HttpFrameworkType
	GetDatabaseType() DatabaseType

	GetMySqlBookDatabase() string
	GetMySqlBookUser() string
	GetMySqlBookPassword() string
	GetMySqlBookHost() string
	GetMySqlBookPort() int

	GetPostgresBookDatabase() string
	GetPostgresBookUser() string
	GetPostgresBookPassword() string
	GetPostgresBookHost() string
	GetPostgresBookPort() int
}
