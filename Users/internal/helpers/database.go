package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func OpenDB() (*sql.DB, error) { // ferme et ouvre la base de donnée 
	db, err := sql.Open("sqlite3", "file:users.db")
	if err != nil {
		db.SetMaxOpenConns(1)
	}
	return db, err
}
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logrus.Errorf("error closing db : %s", err.Error())
	}
}
