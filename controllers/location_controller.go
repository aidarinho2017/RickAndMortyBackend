package controllers

import (
	"RickAndMortyBackend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	apiURL := fmt.Sprintf("%slocation", baseUrl)
	log.Printf("Fetching locations from: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching locations: %v, Status code: %d", err, resp.StatusCode)
		http.Error(w, "Failed to get locations", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var locations models.AllLocations
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		log.Printf("Error decoding locations: %v", err)
		http.Error(w, "Error decoding locations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		log.Printf("Error encoding locations json: %v", err)
		http.Error(w, "Error encoding locations", http.StatusInternalServerError)
		return
	}
	log.Println("Successfully fetched locations")
}

func GetLocationByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/locations/"):]
	if id == "" {
		log.Println("Error: Missing ID")
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("%slocation/%s", baseUrl, id)
	log.Printf("Fetching location by id from: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching location: %v, Status code: %d", err, resp.StatusCode)
		http.Error(w, "Failed to get location", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var location models.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Printf("Error decoding location: %v", err)
		http.Error(w, "Error decoding location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(location); err != nil {
		log.Printf("Error encoding location json: %v", err)
		http.Error(w, "Error encoding location", http.StatusInternalServerError)
		return
	}
	log.Println("Successfully fetched location by ID")
}
