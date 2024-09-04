package database

import (
	"config"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB(conf *config.Config) {
	dbConf := conf.Database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Printf("Connected to database %s", dbConf.Database)

	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func CountTableRows(table string) (int, error) {
	statement, err := Db.Prepare("SELECT COUNT(*) FROM " + table)
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow()

	var count int
	err = row.Scan(&count)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Print(err)
		}
		return 0, err
	}

	return count, nil
}
