package controllers

import (
	"RickAndMortyBackend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const baseUrl = "https://rickandmortyapi.com/api/"

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	query := url.Values{}
	for key, val := range params {
		query.Set(key, val[0])
	}
	apiURL := fmt.Sprintf("%scharacter/?%s", baseUrl, query.Encode())

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get characters", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var data models.AllCharacters
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding character data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func GetCharacterByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/character/"):]
	if id == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("%scharacter/%s", baseUrl, id)

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get character", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var character models.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		http.Error(w, "Error decoding character", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}
