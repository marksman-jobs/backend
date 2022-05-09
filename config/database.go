package config

import (
	"database/sql"

	"github.com/marksman-jobs/backend/db"
)

type Databases struct {
	Postgres *db.Queries
}

func InitializeDatabase(configuration Config) (*Databases, error) {

	postgresDriver, err := sql.Open("postgres", configuration.POSTGRES_URI)
	if err != nil {
		return nil, err
	}

	postgres := db.New(postgresDriver)

	database := &Databases{
		Postgres: postgres,
	}

	return database, nil

}
