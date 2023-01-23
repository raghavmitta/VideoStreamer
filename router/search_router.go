package router

import (
	"VideoStreamer/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SearchResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	fmt.Println(query)
	log.Println("Exact search for query{}", query)
	pages := service.ExactSearch(query)
	if pages == nil || len(pages) == 0 {
		log.Println("No result found")
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Println("Total pages:{} found for query ", len(pages), query)
		json.NewEncoder(w).Encode(pages)
	}
	return
}
