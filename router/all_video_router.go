package router

import (
	"VideoStreamer/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	pages := service.GetAllData()
	log.Println("Getting all the videos")
	if pages == nil || len(pages) == 0 {
		log.Println("No result found")
		http.NotFound(w, r)
	} else {
		json.NewEncoder(w).Encode(pages)
	}
	return
}
