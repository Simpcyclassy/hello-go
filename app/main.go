package main

import (
	"net/http"

	"github.com/Simpcyclassy/hello-go/app/cache"
	"github.com/Simpcyclassy/hello-go/app/config"
	"github.com/Simpcyclassy/hello-go/app/handlers"
	"github.com/rs/zerolog/log"
)

func main() {
	fileServer := http.FileServer(http.Dir("../client"))
	http.Handle("/", fileServer)

	inMemCache := cache.NewLRUCache(
		config.Config.Memcache.TTL,
		config.Config.Memcache.Size,
		config.Config.Memcache.Prunesize,
	)
	log.Debug().Str("memcache name", inMemCache.GetName()).Msg("memcache name in main.go")
	informationHandler := handlers.New(inMemCache)
	http.HandleFunc("/tree", handlers.FavoriteTreeHandler)
	http.Handle("/info", informationHandler)

	log.Info().Msg("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal().Err(err)
	}
}
