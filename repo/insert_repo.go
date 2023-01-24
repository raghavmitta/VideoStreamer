package repo

import (
	"VideoStreamer/db"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

func InsertData(result *youtube.SearchResult) {
	insertQuery := "REPLACE INTO video VALUES ( ?,?,?,?,?,?,?)"
	db := db.ConnectDB()
	_, err := db.Exec(insertQuery, result.Id.VideoId, timeParser(result.Snippet.PublishedAt), result.Snippet.Title, result.Snippet.Description, result.Snippet.Thumbnails.Default.Url, result.Snippet.Thumbnails.Medium.Url, result.Snippet.Thumbnails.High.Url)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("Recorded with VideoID: ", result.Id.VideoId, " inserted")
	}
	defer db.Close()
	return

}

func timeParser(input string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, input)
	if err != nil {
		log.Println(err)
		return timestamp

	} else {
		return timestamp
	}

}
