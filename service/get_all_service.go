package service

import (
	"VideoStreamer/model"
	"VideoStreamer/repo"
	"VideoStreamer/util"
)

func GetAllData() []model.Page {

	return util.GetPaginated(repo.FetchAllData())
}
