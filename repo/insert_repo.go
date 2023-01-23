package repo

import (
	"VideoStreamer/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

func InsertData(result *youtube.SearchResult) {
	insertQuery := "REPLACE INTO video VALUES ( ?,?,?,?,?,?,?)"
	db := config.ConnectDB()
	_, err := db.Exec(insertQuery, result.Id.VideoId, timeParser(result.Snippet.PublishedAt), result.Snippet.Title, result.Snippet.Description, result.Snippet.Thumbnails.Default.Url, result.Snippet.Thumbnails.Medium.Url, result.Snippet.Thumbnails.High.Url)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("Recorded with VideoID: {} inserted", result.Id)
	}
	defer db.Close()
	return

}

func timeParser(input string) time.Time {
	timestamp, err := time.Parse(time.RFC3339, input)
	if err != nil {

		fmt.Println(err)
		return timestamp

	} else {
		fmt.Println(timestamp)
		return timestamp
	}

}
