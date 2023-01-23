package db

import (
	"VideoStreamer/config"
	"database/sql"
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
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + configVar.Database.DbName)
	if err != nil {
		log.Fatalln(err.Error())
		return
	} else {
		log.Println("Successfully created database..")
	}
	_, err = db.Exec("USE " + configVar.Database.DbName)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println("DB selected successfully..")
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS video(video_id varchar(30) NOT NULL, published_at timestamp, title  varchar(120),description  varchar(300),defaultThumbnail  varchar(60),mediumThumbnail  varchar(60),highThumbnail varchar(60), PRIMARY KEY (video_id));")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Table created successfully..")
		return
	}
	defer db.Close()
	return

}
