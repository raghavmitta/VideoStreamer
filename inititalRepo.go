package main

import (
	"VideoStreamer/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectDB()
	if db != nil {
		createDatabase(*db)
	}
}
func createDatabase(db sql.DB) {
	/*_, err := db.Exec("CREATE DATABASE testDB")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Successfully created database..")
	}*/
	_, err := db.Exec("USE testDB")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DB selected successfully..")
	stmt, err := db.Prepare("CREATE Table video(video_id varchar(30) NOT NULL, published_at timestamp, title  varchar(120),description  varchar(300),defaultThumbnail  varchar(60),mediumThumbnail  varchar(60),highThumbnail varchar(60), PRIMARY KEY (video_id));")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Table created successfully..")
		return
	}
	defer db.Close()

}
