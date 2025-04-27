package controllers

import (
	"RickAndMortyBackend/models"
	"encoding/json"
	"fmt"
	"log" // Import the log package
	"net/http"
)

const baseUrl = "https://rickandmortyapi.com/api/"

func GetCharacters(w http.ResponseWriter) { // Removed unused 'r'
	apiURL := fmt.Sprintf("%scharacter", baseUrl)
	log.Printf("Fetching characters from: %s", apiURL) // Added log

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching characters: %v, Status Code: %d", err, resp.StatusCode) // Added log
		http.Error(w, "Failed to get characters", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var characters models.AllCharacters
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		log.Printf("Error decoding characters: %v", err) // Added log
		http.Error(w, "Error decoding characters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(characters); err != nil {
		log.Printf("Error encoding characters to JSON: %v", err)             // Added log
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) //Added error
		return
	}
	log.Println("Successfully fetched and sent characters") // Added log
}

func GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/characters/"):] // Corrected path to "/characters/" to match main.go
	if id == "" {
		log.Println("Error: Missing ID") // Added log
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	apiURL := fmt.Sprintf("%scharacter/%s", baseUrl, id)
	log.Printf("Fetching character by ID from: %s", apiURL) // Added log

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching character by ID: %v, Status Code: %d", err, resp.StatusCode) // Added log
		http.Error(w, "Failed to get character", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var character models.Character
	if err := json.NewDecoder(resp.Body).Decode(&character); err != nil {
		log.Printf("Error decoding character: %v", err) // Added log
		http.Error(w, "Error decoding character", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(character); err != nil { //added error handling
		log.Printf("Error encoding character to JSON: %v", err)
		http.Error(w, "Error encoding character JSON", http.StatusInternalServerError)
		return
	}
	log.Println("Successfully fetched and sent character by ID") // Added log
}
