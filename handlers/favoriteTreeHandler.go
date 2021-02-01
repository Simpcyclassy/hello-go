package handlers

import (
	"net/http"
)

// FavoriteTreeHandler handles favorite query
func FavoriteTreeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tree" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request, please go to /tree"))
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Method not found."))
		return
	}

	query := r.URL.Query()
	favoriteTree := query.Get("favoriteTree")

	if favoriteTree == "" {
		w.WriteHeader(200)
		w.Write([]byte("Please tell me your favorite tree"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("It's nice to know that your favorite tree is a " + favoriteTree))
}
