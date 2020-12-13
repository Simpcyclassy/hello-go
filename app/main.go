package app

import (
	"log"
	"net/http"
)

func favoriteTreeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request, please go to /"))
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

func main() {
	http.HandleFunc("/", favoriteTreeHandler)

	log.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
