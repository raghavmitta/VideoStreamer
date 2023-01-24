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
	log.Println("Exact search for query", query)
	pages := service.ExactSearch(query)
	if pages == nil || len(pages) == 0 {
		log.Println("No result found for Query:", query)
		http.NotFound(w, r)
	} else {
		log.Println("Total pages:", len(pages), "found for Query", query)
		json.NewEncoder(w).Encode(pages)
	}
	return
}
func PartialSearchResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	log.Println("Partial search for query", query)
	pages := service.PartialSearch(query)
	if pages == nil || len(pages) == 0 {
		log.Println("No result found for Query:", query)
		http.NotFound(w, r)
	} else {
		log.Println("Total pages:", len(pages), "found for query", query)
		json.NewEncoder(w).Encode(pages)
	}
	return
}
