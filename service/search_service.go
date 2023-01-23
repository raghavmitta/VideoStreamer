package service

import (
	"VideoStreamer/model"
	"VideoStreamer/repo"
	"VideoStreamer/util"
)

func ExactSearch(key string) []model.Page {
	return util.GetPaginated(repo.SearchData(key))
}
