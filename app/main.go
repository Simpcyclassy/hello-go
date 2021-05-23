package main

import (
	"log"
	"net/http"

	"github.com/Simpcyclassy/hello-go/app/handlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("../client"))
	http.Handle("/", fileServer)

	http.HandleFunc("/tree", handlers.FavoriteTreeHandler)
	http.HandleFunc("/info", handlers.InformationHandler)

	log.Println("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
