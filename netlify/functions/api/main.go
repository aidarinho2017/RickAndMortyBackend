package main

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"net/http"
	"strings"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()

	mux.Handle("/characters", http.HandlerFunc(handleCharacters))
	mux.Handle("/characters/", http.HandlerFunc(handleCharacters))

	mux.Handle("/locations", http.HandlerFunc(handleLocations))
	mux.Handle("/locations/", http.HandlerFunc(handleLocations))

	mux.Handle("/episodes", http.HandlerFunc(handleEpisodes))
	mux.Handle("/episodes/", http.HandlerFunc(handleEpisodes))
}

func handleCharacters(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/characters" {
		controllers.GetCharacters(w, r)
	} else if strings.HasPrefix(path, "/characters/") {
		controllers.GetCharacterByID(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func handleLocations(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/locations" {
		controllers.GetLocations(w, r)
	} else if strings.HasPrefix(path, "/locations/") {
		controllers.GetLocationByID(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func handleEpisodes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/episodes" {
		controllers.GetEpisodes(w, r)
	} else if strings.HasPrefix(path, "/episodes/") {
		controllers.GetEpisodeByID(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// Correct Netlify serverless Go function handler
func Handler(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(mux).ServeHTTP(w, r)
}

func main() {}
