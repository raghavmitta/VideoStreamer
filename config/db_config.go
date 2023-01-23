package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:Moto@688729@tcp(127.0.0.1:3306)/testDb?parseTime=true")

	if err != nil {
		panic(err.Error())
		return nil
	}
	return db
}
