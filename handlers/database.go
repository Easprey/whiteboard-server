package handlers

import (
	"database/sql"
	"log"
)

type DataBase struct {
	ConnString string `long:"database" short:"d" description:"MySQL connection string of the form [username[:password]@[protocol[(address)]]/dbname" env:"DB_CONN"`
	*sql.DB           // anonymous so methods are embedded
}

func (db *DataBase) Init() {
	// fail if ConnString not set
	if db.ConnString == "" {
		log.Panic("No database connection string provided")
	}

	// attempt to open database connection
	var err error
	db.DB, err = sql.Open("mysql", db.ConnString)
	if err != nil {
		log.Panic(err)
	}

	// attempt to contact database
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}
