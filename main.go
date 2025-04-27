package main

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"log"
	"net/http"
	"os"
	"strings" // Keep if used in controllers or handler logic below
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}

	mux := http.NewServeMux()

	// --- Define Handlers (without EnableCORS wrapper here) ---

	// Handler for character routes
	charactersHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example routing logic: (Adjust based on your controller needs)
		// Recommend trimming slash for consistent matching
		path := strings.TrimSuffix(r.URL.Path, "/")
		parts := strings.Split(path, "/") // e.g., ["", "characters", "1"]

		if len(parts) == 3 && parts[1] == "characters" && parts[2] != "" {
			log.Printf("Routing to GetCharacterByID for path: %s", r.URL.Path)
			controllers.GetCharacterByID(w, r) // Assumes this extracts ID from r.URL.Path
		} else if path == "/characters" {
			log.Printf("Routing to GetCharacters for path: %s", r.URL.Path)
			controllers.GetCharacters(w, r) // Corrected:  GetCharacters now takes only w
		} else {
			http.NotFound(w, r)
		}
	})

	// Handler for location routes (Adapt logic similarly if needed)
	locationsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Simplified if controllers handle sub-paths
		path := strings.TrimSuffix(r.URL.Path, "/")
		if strings.HasPrefix(path, "/locations/") && path != "/locations" {
			controllers.GetLocationByID(w, r)
		} else if path == "/locations" {
			controllers.GetLocations(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Handler for episode routes (Adapt logic similarly if needed)
	episodesHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.URL.Path, "/")
		if strings.HasPrefix(path, "/episodes/") && path != "/episodes" {
			controllers.GetEpisodeByID(w, r)
		} else if path == "/episodes" {
			controllers.GetEpisodes(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// --- Register Handlers with ServeMux ---
	// Use patterns that make sense for your routing logic.
	// Handling both "/path" and "/path/" can be safer.
	mux.Handle("/characters", charactersHandler)
	mux.Handle("/characters/", charactersHandler) // Matches /characters/ and /characters/id etc.

	mux.Handle("/locations", locationsHandler)
	mux.Handle("/locations/", locationsHandler)

	mux.Handle("/episodes", episodesHandler)
	mux.Handle("/episodes/", episodesHandler)

	// --- Apply CORS Middleware Globally ---
	// Wrap the entire mux with the CORS middleware
	corsEnabledMux := middleware.EnableCORS(mux)

	log.Println("Server starting on port", port)

	// --- Start Server ---
	// Listen using the mux that has the CORS middleware applied globally
	log.Fatal(http.ListenAndServe(":"+port, corsEnabledMux)) // Use corsEnabledMux here
}
