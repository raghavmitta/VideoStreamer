package handler

import (
	"VideoStreamer/router"
	"github.com/gorilla/mux"
)

func Handle() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", router.GetAllVideos)
	myRouter.HandleFunc("/search/{query}", router.SearchResult) //endPoint to search data with given query

}
