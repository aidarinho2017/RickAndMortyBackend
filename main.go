package main

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/characters", middleware.EnableCORS(http.HandlerFunc(controllers.GetCharacters)))
	mux.Handle("/characters/{id}", middleware.EnableCORS(http.HandlerFunc(controllers.GetCharacterByID)))

	mux.Handle("/locations", middleware.EnableCORS(http.HandlerFunc(controllers.GetLocations)))
	mux.Handle("/locations/{id}", middleware.EnableCORS(http.HandlerFunc(controllers.GetLocationByID)))

	mux.Handle("/episodes", middleware.EnableCORS(http.HandlerFunc(controllers.GetEpisodes)))
	mux.Handle("/episodes/{id}", middleware.EnableCORS(http.HandlerFunc(controllers.GetEpisodeByID)))

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
