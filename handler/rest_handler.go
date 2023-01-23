package handler

import (
	"VideoStreamer/config"
	"VideoStreamer/router"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Handle() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", router.GetAllVideos)
	myRouter.HandleFunc("/search/{query}", router.SearchResult) //endPoint to search data with given query
	myRouter.HandleFunc("/partial-search/{query}", router.PartialSearchResult)
	log.Fatal(http.ListenAndServe(":"+config.GetConfig().Server.Port, myRouter))

}
