package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// db is a package global variable
var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}
