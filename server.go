package main

import (
	"log"
	"net/http"
	"github.com/Simpcyclassy/hello-go/handlers/favoriteTreeHandler"
	"github.com/Simpcyclassy/hello-go/handlers/weatherHandler"
)

func main() {
	fileServer := http.FileServer(http.Dir("../client"))
	http.Handle("/", fileServer)

	http.HandleFunc("/tree", favoriteTreeHandler)
	http.HandleFunc("/weather", weatherHandler)

	log.Println("Starting server at port 8000")
	err := http.ListenAndServe(":8000", nil);
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

