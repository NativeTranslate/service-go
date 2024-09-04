package database

import (
	"config"
	"database/sql"
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
