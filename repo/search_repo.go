package repo

import (
	"VideoStreamer/db"
	"database/sql"
	"log"
)

func SearchData(key string) *sql.Rows {
	db := db.ConnectDB()
	key = "%" + key + "%"
	results, err := db.Query("SELECT * FROM video where ((title like ?) or (description like ?)) order by published_at desc;", key, key)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	db.Close()
	return results
}
