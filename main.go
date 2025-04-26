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

	mux.Handle("/characters/", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/characters" || r.URL.Path == "/characters/" {
			controllers.GetCharacters(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/characters/") {
			controllers.GetCharacterByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})))

	mux.Handle("/locations/", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/locations" || r.URL.Path == "/locations/" {
			controllers.GetLocations(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/locations/") {
			controllers.GetLocationByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})))

	mux.Handle("/episodes/", middleware.EnableCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/episodes" || r.URL.Path == "/episodes/" {
			controllers.GetEpisodes(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/episodes/") {
			controllers.GetEpisodeByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})))
}

// Exported handler for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
