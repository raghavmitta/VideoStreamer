package util

import (
	"VideoStreamer/config"
	"VideoStreamer/model"
	"database/sql"
	"fmt"
)

var pageSize int

func GetResult(results *sql.Rows) model.Video { // converts sql data into Video model
	var video model.Video
	err := results.Scan(&video.VideoId, &video.PublishedAt, &video.Title, &video.Description, &video.DefaultThumbnail, &video.MediumThumbnail, &video.HighThumbnail)
	if err != nil {
		fmt.Println(err.Error())
	}
	return video
}
func GetPaginated(results *sql.Rows) []model.Page {
	pageSize = config.GetConfig().Pagination.PageSize
	var pages = make([]model.Page, 0, pageSize) //initial no.of pages=10
	var pageNo = 0
	tempPage := model.NewPage(pageNo, pageSize) //pageSize=10
	for results.Next() {
		tempPage.Results = append(tempPage.Results, GetResult(results))
		if len(tempPage.Results) == tempPage.Size {
			pages = append(pages, *tempPage)           //save current page in pages
			pageNo++                                   // increase page
			tempPage = model.NewPage(pageNo, pageSize) //create new page for left data
		}
	}
	if len(tempPage.Results) != 0 {
		pages = append(pages, *tempPage)
	}
	return pages
}
func PaginatedSearch(keys []string, videoIdMapper map[string]model.Video) []model.Page {
	pageSize = config.GetConfig().Pagination.PageSize
	var pages = make([]model.Page, 0, pageSize)
	var pageNo = 0
	tempPage := model.NewPage(pageNo, pageSize)
	for _, key := range keys {
		tempPage.Results = append(tempPage.Results, videoIdMapper[key])
		if len(tempPage.Results) == tempPage.Size {
			pages = append(pages, *tempPage)
			pageNo++
			tempPage = model.NewPage(pageNo, pageSize)
		}
	}
	if len(tempPage.Results) != 0 {
		pages = append(pages, *tempPage)
	}
	return pages
}
