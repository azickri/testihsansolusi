package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	User         string
	DatabaseName string
	Password     string
	Host         string
	Client       *sql.DB
}

func (postgreSQL *PostgreSQL) Connect() {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		postgreSQL.Host, postgreSQL.User, postgreSQL.Password, postgreSQL.DatabaseName,
	)

	client, err := sql.Open("postgres", config)
	if err != nil {
		fmt.Println("CRITICAL: error connect to postgresql", err.Error())
	}

	postgreSQL.Client = client
}
