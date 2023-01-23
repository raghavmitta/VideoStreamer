package service

import (
	"VideoStreamer/repo"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

var publisherTime string

func StartService(duration time.Duration) {
	log.Println("Consuming Youtube V3 Api")
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyDvu6tqmK0nZvwMLZ_mddV4OQWWbW5g7lo")) //connection with given API Key
	ticker := time.NewTicker(duration * time.Second)                                                             //initialize ticker for async update
	if err != nil {
		log.Println("Failure to connect")
		log.Println(err.Error())
		return
	}
	publisherTime = time.Now().Add(-1 * time.Hour).Format(time.RFC3339) //fetching results after currentTime-1hour
	for _ = range ticker.C {
		go searchByKey(youtubeService.Search) // Async consumption using goroutine
	}
	return
}
func searchByKey(service *youtube.SearchService) {
	//call to Youtube  API with query:tutorial, orderby Date,language:English,publishAfter:given time
	call := service.List([]string{"snippet"}).MaxResults(50).Q("tutorial").Order("date").Type("video").RelevanceLanguage("en").PublishedAfter(publisherTime)
	//Update publish time to fetch latest result in last call
	publisherTime = time.Now().Format(time.RFC3339)
	response, err := call.Do()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("No. of new Result fetch" + string(response.PageInfo.ResultsPerPage))
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			//save data to db
			repo.InsertData(item)
		}

	}

	return
}
