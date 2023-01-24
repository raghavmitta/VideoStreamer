package service

import (
	"VideoStreamer/config"
	"VideoStreamer/repo"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

var publisherTime string
var keys []string
var currentKeyIndex int

func StartService(duration time.Duration) {
	currentKeyIndex = 0
	keys = config.GetConfig().Api.Keys
	ticker := time.NewTicker(duration * time.Second)                    //initialize ticker for async update
	publisherTime = time.Now().Add(-1 * time.Hour).Format(time.RFC3339) //fetching results after currentTime-1hour
	go multipleKeySupport()
	for _ = range ticker.C {
		if len(keys) <= currentKeyIndex {
			log.Println("Stopping ticker,all the keys exhausted ")
			break
		} else {
			go multipleKeySupport() // Async consumption using goroutine
		}
	}
	return
}
func multipleKeySupport() {
	if len(keys) == 0 {
		log.Fatalln("Please Provide Api Keys in config")
	}
	if len(keys) <= currentKeyIndex {
		return
	}
	log.Println("Consuming Youtube V3 Api to fetch results after ", publisherTime)
	youtubeService := createService(keys[currentKeyIndex])
	searchByKey(youtubeService.Search)
	return

}
func createService(key string) *youtube.Service {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(key))
	if err != nil {
		log.Println("Error while creating Youtube service")
		log.Fatalln(err.Error())
	}
	return youtubeService
}

func searchByKey(service *youtube.SearchService) {
	//call to Youtube  API with query:tutorial, orderby Date,language:English,publishAfter:given time
	call := service.List([]string{"snippet"}).MaxResults(50).Q(config.GetConfig().Api.Query).Order("date").Type("video").RelevanceLanguage("en").PublishedAfter(publisherTime)
	response, err := call.Do()
	if err != nil {
		currentKeyIndex++
		log.Println("Current Key exhausted, trying next key")
		log.Println(err.Error())
		return
	} else {
		publisherTime = time.Now().Format(time.RFC3339) //Update publish time to fetch latest result in last call
	}
	if response.PageInfo.ResultsPerPage > 0 {
		log.Println("Adding ", response.PageInfo.ResultsPerPage, " new results to db")
	}
	//log.Println("No. of new Result fetch" + string(response.PageInfo.ResultsPerPage))
	for _, item := range response.Items {
		//save data to db
		repo.InsertData(item)

	}

	return
}
