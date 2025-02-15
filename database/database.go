package database

import (
	"database/sql"
	"log"
	"time"
)

var Dbconn *sql.DB

//SetupDatabase creates a connection to the database

func SetupDatabase() {
	var err error
	Dbconn, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	Dbconn.SetMaxOpenConns(3)
	Dbconn.SetMaxIdleConns(3)
	Dbconn.SetConnMaxLifetime(60 * time.Second)
}
