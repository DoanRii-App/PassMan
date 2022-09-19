package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func fDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./passman.db")
	st, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (User varchar(255), Pass varchar(255), Email varchar(255), Otp varchar(255), Tag varchar(255), Oki varchar(255))")
	st.Exec()
	return
}