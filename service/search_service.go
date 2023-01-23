package service

import (
	"VideoStreamer/model"
	"VideoStreamer/repo"
	"VideoStreamer/util"
	"sort"
	"strings"
)

func ExactSearch(key string) []model.Page {
	return util.GetPaginated(repo.SearchData(key))
}

// This function will return video sorted in desceding order of no.words in query it title or description contains and publishing time
func PartialSearch(key string) []model.Page {
	var keyMapper = make(map[string]int)             //Map to hold Video id and frequency of words its title and description contains
	var videoIdMapper = make(map[string]model.Video) //Map to hold Video id and Video object
	split := strings.Split(key, " ")                 //convert query into collection of words
	for _, searchKey := range split {
		if len(searchKey) > 1 { // ignore one lettered word a,v,d,c,e
			videoArray := getPartialResponse(searchKey) //get data in form of []Video
			for _, video := range videoArray {
				videoIdMapper[video.VideoId] = video         //populate VideoId mapper with videoId and Video object
				value, isPresent := keyMapper[video.VideoId] //check if VideoID is already present in KeyMapper
				if isPresent {
					keyMapper[video.VideoId] = value + 1 //if yes, increase the frequency
				} else {
					keyMapper[video.VideoId] = 1 // if no, store it with frequency one
				}

			}
		}
	}
	var keys []string
	//create a slice of video id present in KeyMapper
	for mapKey := range keyMapper {
		keys = append(keys, mapKey)
	}
	if len(keys) == 0 {
		return nil
	}
	//Sort the videoIds in decreasing order of frequency of words it matched with and
	//if videoId's frequencies are same then sort in order of their publishing time
	sort.SliceStable(keys, func(i, j int) bool {
		if keyMapper[keys[i]] == keyMapper[keys[j]] {
			return videoIdMapper[keys[i]].PublishedAt.After(videoIdMapper[keys[j]].PublishedAt)
		} else {
			return keyMapper[keys[i]] > keyMapper[keys[j]]
		}
	})
	//Convert result into paginated form
	return util.PaginatedSearch(keys, videoIdMapper)
}
func getPartialResponse(key string) []model.Video {
	videoArray := make([]model.Video, 0, 10)
	result := repo.SearchData(key) // search in repo for given word
	//covert sql data into []Video
	for result.Next() {
		video := util.GetResult(result)
		videoArray = append(videoArray, video)
	}
	return videoArray

}
