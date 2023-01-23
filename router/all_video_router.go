package router

import (
	"VideoStreamer/service"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pages := service.GetAllData()
	log.Println("Getting all the videos")
	if len(pages) == 0 {
		log.Println("No result found")
		w.WriteHeader(http.StatusNoContent)
	} else {
		json.NewEncoder(w).Encode(pages)
	}
	return
}
