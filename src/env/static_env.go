package env

type StaticEnv struct{}

func (se StaticEnv) GetHttpFramework() HttpFrameworkType {
	return HttpFrameworkType("FIBER")
}

func (se StaticEnv) GetDatabaseType() DatabaseType {
	return DatabaseType("MYSQL")
}

func (se StaticEnv) GetMySqlBookDatabase() string {
	return "book"
}

func (se StaticEnv) GetMySqlBookUser() string {
	return "book"
}

func (se StaticEnv) GetMySqlBookPassword() string {
	return "book"
}

func (se StaticEnv) GetMySqlBookHost() string {
	return "localhost"
}

func (se StaticEnv) GetMySqlBookPort() int {
	return 3306
}

func (se StaticEnv) GetPostgresBookDatabase() string {
	return "book"
}

func (se StaticEnv) GetPostgresBookUser() string {
	return "book"
}

func (se StaticEnv) GetPostgresBookPassword() string {
	return "book"
}

func (se StaticEnv) GetPostgresBookHost() string {
	return "localhost"
}

func (se StaticEnv) GetPostgresBookPort() int {
	return 5432
}
