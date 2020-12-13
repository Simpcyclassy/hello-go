package main

import (
	"log"
	"net/http"
)

func favoriteTreeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Bad request, please go to /", http.StatusBadRequest)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 Method not found.", http.StatusNotFound) //StatusMethodNotAllowed = 405 // RFC 7231, 6.5.5 or  http.NotFound(w, r)
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

func main() {
	http.HandleFunc("/", favoriteTreeHandler)

	log.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
