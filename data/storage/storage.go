package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func New() {
	db, err := sql.Open(os.Getenv("SQL_DRIVER"), os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
