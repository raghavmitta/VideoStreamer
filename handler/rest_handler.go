package handler

import (
	"VideoStreamer/router"
	"github.com/gorilla/mux"
)

func Handle() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", router.GetAllVideos) //endPoint to fetch all data
}
