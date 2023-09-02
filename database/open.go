package database

import (
	"gocron-sample/database/errors"
	"gocron-sample/database/query"
	"log"

	"github.com/jmoiron/sqlx"
)

var dbContext *sqlx.DB

func OpenDb() {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	dbContext = db
}

func getDbContext() (*sqlx.DB, error) {
	if dbContext == nil {
		return nil, errors.ErrDbContextNotInitialized
	}
	return dbContext, nil
}

// Function for creating basic tables.
func CreateTables() {
	if ctx, err := getDbContext(); err != nil {
		ctx.MustExec(query.CreateTables)
	} else {
		log.Fatal(err)
	}
}
