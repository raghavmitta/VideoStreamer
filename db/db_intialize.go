package db

import (
	"VideoStreamer/config"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB
var configVar *config.Config

func ConnectDB() *sql.DB {
	var err error
	configVar = config.GetConfig()
	db, err = sql.Open(configVar.Database.DriverName, configVar.Database.HostName+"/"+configVar.Database.DbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
		return nil
	}
	return db
}

func Initialize() {
	ConnectDB()
	queries := []string{
		"CREATE DATABASE IF NOT EXISTS " + configVar.Database.DbName,
		"USE " + configVar.Database.DbName,
		"CREATE TABLE IF NOT EXISTS video(video_id varchar(30) NOT NULL, published_at timestamp, title  varchar(120),description  varchar(300),defaultThumbnail  varchar(60),mediumThumbnail  varchar(60),highThumbnail varchar(60), PRIMARY KEY (video_id),INDEX (title,description))"}
	fmt.Print(queries[1])
	for _, query := range queries {
		_, err := db.Query(query)
		handleError(err)
	}
	log.Println("Database successfully initialised")
	defer db.Close()
	return

}
func handleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}

}
