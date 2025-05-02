package main

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

func routeHandler(basePath string, listHandler, detailHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.URL.Path, "/")
		if path == basePath {
			listHandler(w, r)
		} else if strings.HasPrefix(path, basePath+"/") {
			detailHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Register all handlers in one line each
	mux.Handle("/characters", routeHandler("/characters", controllers.GetCharacters, controllers.GetCharacterByID))
	mux.Handle("/characters/", routeHandler("/characters", controllers.GetCharacters, controllers.GetCharacterByID))

	mux.Handle("/locations", routeHandler("/locations", controllers.GetLocations, controllers.GetLocationByID))
	mux.Handle("/locations/", routeHandler("/locations", controllers.GetLocations, controllers.GetLocationByID))

	mux.Handle("/episodes", routeHandler("/episodes", controllers.GetEpisodes, controllers.GetEpisodeByID))
	mux.Handle("/episodes/", routeHandler("/episodes", controllers.GetEpisodes, controllers.GetEpisodeByID))

	// Apply CORS
	corsEnabledMux := middleware.EnableCORS(mux)

	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, corsEnabledMux))
}
