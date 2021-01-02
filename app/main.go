package main

import (
	"log"
	"net/http"
)

// struct type to capture our data
// type WeatherCollection struct {
// 	tempperatureValues Temperature
// 	windValues Wind
// }

// type Temperature struct {
// 	celcius int32
// 	farenheit int32
// }

// type Wind struct {
// 	speed int32
// 	direction int32
// }

// func that makes a request to the API (it's probably a JSON body)
// unmarshel the JSON to our custom struct

// New handler for /weather (only serve on this path), only take get requests.
// It should return a 200 with a json body of the WeatherCollection Data.

// And we should do tests ;)

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
	// fileServer := http.FileServer(http.Dir("../static"))
	// http.Handle("/", fileServer)

	http.HandleFunc("/", favoriteTreeHandler)

	log.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
