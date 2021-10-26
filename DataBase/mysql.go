package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/camdb?parseTime=true")
	if err != nil {
		log.Printf("SQL Error : %s", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Ping Error : %s", err.Error())
		return
	}
}
