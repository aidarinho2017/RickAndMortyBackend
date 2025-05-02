// api/index.go
package handler

import (
	"RickAndMortyBackend/controllers"
	"RickAndMortyBackend/middleware"
	"net/http"
	"strings"
)

// Handler is the Vercel entrypoint for any /api/* request.
func Handler(w http.ResponseWriter, r *http.Request) {
	// 1) Strip "/api" prefix so our ServeMux sees "/characters", "/episodes/123", etc.
	path := r.URL.Path
	if strings.HasPrefix(path, "/api") {
		path = strings.TrimPrefix(path, "/api")
	}
	r2 := r.Clone(r.Context())
	r2.URL.Path = path

	// 2) Build a ServeMux and register all routes
	mux := http.NewServeMux()

	// Health-check at "/" (i.e. /api)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Rick and Morty API is up 🚀"))
	})

	// List endpoints
	mux.HandleFunc("/characters", controllers.GetCharacters)
	mux.HandleFunc("/locations", controllers.GetLocations)
	mux.HandleFunc("/episodes", controllers.GetEpisodes)

	// Detail endpoints: extract the ID and rewrite it into r.URL.Query()
	mux.HandleFunc("/characters/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/characters/")
		q := r.URL.Query()
		q.Set("id", id)
		r = r.Clone(r.Context())
		r.URL.RawQuery = q.Encode()
		controllers.GetCharacterByID(w, r)
	})
	mux.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/locations/")
		q := r.URL.Query()
		q.Set("id", id)
		r = r.Clone(r.Context())
		r.URL.RawQuery = q.Encode()
		controllers.GetLocationByID(w, r)
	})
	mux.HandleFunc("/episodes/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/episodes/")
		q := r.URL.Query()
		q.Set("id", id)
		r = r.Clone(r.Context())
		r.URL.RawQuery = q.Encode()
		controllers.GetEpisodeByID(w, r)
	})

	// 3) Wrap the mux with CORS middleware and serve
	middleware.EnableCORS(mux).ServeHTTP(w, r2)
}
