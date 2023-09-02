package database

import (
	"gocron-sample/database/errors"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var dbContext *sqlx.DB

func OpenDb(connString string) {
	db, err := sqlx.Open("pgx", connString)
	if err != nil {
		log.Fatal(err)
	}
	dbContext = db
}

func GetDbContext() (*sqlx.DB, error) {
	if dbContext == nil {
		return nil, errors.ErrDbContextNotInitialized
	}
	return dbContext, nil
}
