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
		port = "8080"
	}

	mux := http.NewServeMux()

	// --- Define Handlers ---

	charactersHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.URL.Path, "/")

		// Correct handling: /characters and /characters/{id}
		if path == "/characters" {
			log.Printf("Routing to GetCharacters (list) for path: %s", r.URL.Path)
			controllers.GetCharacters(w, r)
		} else if strings.HasPrefix(path, "/characters/") {
			log.Printf("Routing to GetCharacterByID for path: %s", r.URL.Path)
			controllers.GetCharacterByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	locationsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.URL.Path, "/")

		if path == "/locations" {
			controllers.GetLocations(w, r)
		} else if strings.HasPrefix(path, "/locations/") {
			controllers.GetLocationByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	episodesHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.URL.Path, "/")

		if path == "/episodes" {
			controllers.GetEpisodes(w, r)
		} else if strings.HasPrefix(path, "/episodes/") {
			controllers.GetEpisodeByID(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// --- Register Handlers ---

	mux.Handle("/characters", charactersHandler)
	mux.Handle("/characters/", charactersHandler)

	mux.Handle("/locations", locationsHandler)
	mux.Handle("/locations/", locationsHandler)

	mux.Handle("/episodes", episodesHandler)
	mux.Handle("/episodes/", episodesHandler)

	// --- Wrap everything with CORS middleware ---

	corsEnabledMux := middleware.EnableCORS(mux)

	log.Println("Server starting on port", port)

	log.Fatal(http.ListenAndServe(":"+port, corsEnabledMux))
}
