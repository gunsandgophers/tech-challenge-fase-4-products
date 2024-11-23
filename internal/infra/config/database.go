package config

var (
	DB_HOST     = GetEnv("DB_HOST", "dbproducts")
	DB_POST     = GetEnv("DB_PORT", "5432")
	DB_USER     = GetEnv("POSTGRES_USER", "tech-challenge-fase")
	DB_PASSWORD = GetEnv("POSTGRES_PASSWORD", "tech-challenge-fase")
	DB_NAME     = GetEnv("POSTGRES_DB", "tech-challenge-fase")
)
