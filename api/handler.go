package handler

import (
	"RickAndMortyBackend/controllers"
	"net/http"
	"strings"
)

// Handler is the main entry point for Vercel’s Go serverless function.
// It implements http.HandlerFunc and is invoked for every API request.
func Handler(w http.ResponseWriter, r *http.Request) {
	routeHandler(w, r)
}

// routeHandler inspects the URL path and dispatches to the appropriate controller.
func routeHandler(w http.ResponseWriter, r *http.Request) {
	// CORS headers (example)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle CORS preflight
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Normalize path: remove the "/api" prefix if present.
	path := r.URL.Path
	if strings.HasPrefix(path, "/api/") {
		path = strings.TrimPrefix(path, "/api")
	}
	path = strings.Trim(path, "/")
	segments := strings.Split(path, "/")

	// Basic routing logic
	if len(segments) == 0 || segments[0] == "" {
		http.NotFound(w, r)
		return
	}
	switch segments[0] {
	case "characters":
		if r.Method == http.MethodGet {
			if len(segments) == 1 {
				// GET /api/characters
				controllers.GetCharacters(w, r)
				return
			} else if len(segments) == 2 {
				// GET /api/characters/:id
				// Pass the id (e.g. via query parameters) before calling handler
				q := r.URL.Query()
				q.Set("id", segments[1])
				r.URL.RawQuery = q.Encode()
				controllers.GetCharacterByID(w, r)
				return
			}
		}
	case "locations":
		if r.Method == http.MethodGet {
			if len(segments) == 1 {
				// GET /api/locations
				controllers.GetLocations(w, r)
				return
			} else if len(segments) == 2 {
				// GET /api/locations/:id
				q := r.URL.Query()
				q.Set("id", segments[1])
				r.URL.RawQuery = q.Encode()
				controllers.GetLocationByID(w, r)
				return
			}
		}
	case "episodes":
		if r.Method == http.MethodGet {
			if len(segments) == 1 {
				// GET /api/episodes
				controllers.GetEpisodes(w, r)
				return
			} else if len(segments) == 2 {
				// GET /api/episodes/:id
				q := r.URL.Query()
				q.Set("id", segments[1])
				r.URL.RawQuery = q.Encode()
				controllers.GetEpisodeByID(w, r)
				return
			}
		}
	}
	// If no route matches:
	http.NotFound(w, r)
}
