package service

import (
	"VideoStreamer/repo"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

var publisherTime string

func StartService(duration time.Duration) {
	log.Println("Consuming Youtube V3 Api")
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyDvu6tqmK0nZvwMLZ_mddV4OQWWbW5g7lo"))
	ticker := time.NewTicker(duration * time.Second)
	if err != nil {
		log.Println("Failure to connect")
		log.Println(err.Error())
		return
	}
	publisherTime = time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
	for _ = range ticker.C {
		go searchByKey(youtubeService.Search)
	}
	return
}
func searchByKey(service *youtube.SearchService) {

	call := service.List([]string{"snippet"}).MaxResults(50).Q("tutorial").Order("date").Type("video").RelevanceLanguage("en").PublishedAfter(publisherTime)
	publisherTime = time.Now().Format(time.RFC3339)
	response, err := call.Do()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(response.PageInfo.ResultsPerPage)
	fmt.Println(response.PageInfo.TotalResults)
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			fmt.Println(item.Snippet.Description)
			repo.InsertData(item)
		}

	}

	return
}
