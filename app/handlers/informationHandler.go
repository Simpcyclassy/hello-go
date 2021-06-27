package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Simpcyclassy/hello-go/app/cache"
	"github.com/rs/zerolog/log"
)

type informationHandler struct {
	memCache cache.Cache
}

func New(c cache.Cache) http.Handler {
	return &informationHandler{memCache: c}
}

// InformationHandler handles weather and covid requests...
func (h *informationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var informationBody []byte
	query := r.URL.Query()
	q := query.Get("q")
	if len(q) < 1 {
		// Add metric for param to see how many people who make reuests without a query
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not passing a country param"))
		return
	}
	log.Debug().Str("memcache name", h.memCache.GetName()).Msg("memcache name in informationHandler.go")
	cacheBody, err := h.memCache.Get(q)
	log.Info().Err(err).Msgf("%v is our query, %v is our cacheBody", q, cacheBody)
	if err != nil || cacheBody == nil {
		log.Debug().Err(err).Msgf("%v direct proxy calls", cacheBody)

		weatherURL := fmt.Sprintf("%s?q=%s&appid=%s", weatherBASE, q, goDotEnvVariable("APP_ID"))
		covidURL := fmt.Sprintf("%s%s", covidBASE, q)

		info, err := checkUrls(weatherURL, covidURL)
		if err != nil {
			// Internal record of what happened
			log.Error().Err(err).Msg("Problem getting the data")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops something went wrong, please try again later"))
			return
		}

		informationBody, err = json.Marshal(info)
		if err != nil {
			// Internal record of what happened
			log.Error().Err(err).Msg("Problem with marshalling")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops something went wrong, please try again later"))
			return
		}

		var cacheData interface{}
		err = json.Unmarshal(informationBody, &cacheData)
		if err != nil {
			// Add metric
			log.Error().Err(err).Msg("Problem getting interface")
		} else {
			log.Debug().Str("memcache name", h.memCache.GetName()).Msg("memcache name in informationHandler.go where we set the memcache")
			err := h.memCache.Set(q, cacheData)
			log.Debug().Err(err).Msgf("%v is our error. %v is our query, %v is the cacheData", err, q, cacheData)
			cacheGet, err := h.memCache.Get(q)
			log.Debug().Msgf("%v is the cached data gotten immediatetly after setting", cacheGet)

			if err != nil {
				// Add metric
				log.Error().Err(err).Msg("Problem setting cache")
			}
		}

	} else {
		log.Info().Msg("Taking cached response")
		informationBody, err = json.Marshal(cacheBody)
		log.Debug().Msgf("%v is the informationBody data gotten immediatetly after taking cached body", informationBody)

		if err != nil {
			// Internal record of what happened
			log.Error().Err(err).Msg("Error encoding to bytes buffer")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops something went wrong, please try again later"))
			return
		}
	}
	//check data exists in the cache
	//if true move to line w.WriteHeader(http.StatusOK)
	//if not the checUrls and marshall information
	w.WriteHeader(http.StatusOK)
	w.Write(informationBody)
}
