package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	driver string
	path   string
}

func NewSQLite() IDataBase {
	return &SQLite{
		driver: "sqlite3",
		path:   os.Getenv("DB_PATH"),
	}
}

func (s *SQLite) Ping() error {
	db, err := sql.Open(s.driver, s.path)
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	return nil
}
