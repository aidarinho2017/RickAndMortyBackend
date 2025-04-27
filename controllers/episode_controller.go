package controllers

import (
	"RickAndMortyBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetEpisodes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := url.Values{}
	for key, val := range params {
		query.Set(key, val[0])
	}
	apiURL := fmt.Sprintf("%sepisode/?%s", baseUrl, query.Encode())

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get episodes", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data models.AllEpisodes
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding episodes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func GetEpisodeByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/episode/"):]
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("%sepisode/%s", baseUrl, id)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get episode", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var episode models.Episode
	if err := json.NewDecoder(resp.Body).Decode(&episode); err != nil {
		http.Error(w, "Error decoding episode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episode)
}
