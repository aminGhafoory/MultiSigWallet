package models

import (
	"database/sql"
	"fmt"
	"log"
)

type PostgresConfig struct {
	DBurl string
}

func Open(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatal("Can not connect to the DataBase")
		return nil, fmt.Errorf("open : %w", err)
	}

	return db, nil
}
