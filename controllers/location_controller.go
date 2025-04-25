package controllers

import (
	"RickAndMortyBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := url.Values{}
	for key, val := range params {
		query.Set(key, val[0])
	}
	apiURL := fmt.Sprintf("%slocation/?%s", baseUrl, query.Encode())

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get locations", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data models.AllLocations
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding location data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func GetLocationByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/location/"):]
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("%slocation/%s", baseUrl, id)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get location", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var location models.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		http.Error(w, "Error decoding location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}
