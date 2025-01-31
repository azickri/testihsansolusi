package main

import (
	"os"
	"testihsansolusi/database"
	"testihsansolusi/web"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	postgresql := initPostgreSQL()
	web.InitRoute(postgresql)
}

func initPostgreSQL() *database.PostgreSQL {
	POSTGRESQL_USER := os.Getenv("POSTGRESQL_USER")
	POSTGRESQL_DATABASE_NAME := os.Getenv("POSTGRESQL_DATABASE_NAME")
	POSTGRESQL_PASSWORD := os.Getenv("POSTGRESQL_PASSWORD")
	POSTGRESQL_HOST := os.Getenv("POSTGRESQL_HOST")

	var postgressql = database.PostgreSQL{
		User:         POSTGRESQL_USER,
		DatabaseName: POSTGRESQL_DATABASE_NAME,
		Password:     POSTGRESQL_PASSWORD,
		Host:         POSTGRESQL_HOST,
	}

	postgressql.Connect()
	postgressql.Migrate()
	return &postgressql
}
