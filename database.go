package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:admin@/desafio")
	//err = db.Ping()
	if err != nil {
		panic(err)
	}
}
