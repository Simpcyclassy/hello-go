package main

import (
	"log"
	"net/http"

	"github.com/Simpcyclassy/hello-go/cache"
	"github.com/Simpcyclassy/hello-go/config"
	"github.com/Simpcyclassy/hello-go/handlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("../client"))
	http.Handle("/", fileServer)

	inMemCache := cache.NewLRUCache(
		config.Config.Memcache.TTL,
		config.Config.Memcache.Size,
		config.Config.Memcache.Prunesize,
	)
	informationHandler := handlers.New(inMemCache)
	http.HandleFunc("/tree", handlers.FavoriteTreeHandler)
	http.Handle("/info", informationHandler)

	log.Println("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
