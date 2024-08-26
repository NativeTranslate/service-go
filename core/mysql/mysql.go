package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1)/nativetranslate")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	Db = db
}

func CloseDB() error {
	return Db.Close()
}
