package database

import (
	"fmt"
	"testihsansolusi/entity/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	User         string
	DatabaseName string
	Password     string
	Host         string
	Client       *gorm.DB
}

func (postgreSQL *PostgreSQL) Connect() {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		postgreSQL.Host, postgreSQL.User, postgreSQL.Password, postgreSQL.DatabaseName,
	)

	client, err := gorm.Open(postgres.Open(config))
	if err != nil {
		fmt.Println("CRITICAL: error connect to postgresql", err.Error())
	}

	postgreSQL.Client = client
}

func (postgresql *PostgreSQL) Migrate() {
	postgresql.Client.AutoMigrate(&model.Account{})
	postgresql.Client.AutoMigrate(&model.AccountHistory{})
}
