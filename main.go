package main

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}

	mux := http.NewServeMux()

	// Characters
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

	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
