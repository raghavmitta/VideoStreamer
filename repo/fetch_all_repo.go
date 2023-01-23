package repo

import (
	"VideoStreamer/config"
	"database/sql"

	"log"
)

func FetchAllData() *sql.Rows {
	db := config.ConnectDB()
	results, err := db.Query("SELECT * FROM video order by published_at desc;")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	db.Close()
	return results

}
