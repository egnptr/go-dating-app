package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

func NewSQLiteRepository() Repo {
	os.Remove("./testing.db")

	db, err := sql.Open("sqlite3", "./testing.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := insertUserTable

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	return &sqliteRepo{}
}
