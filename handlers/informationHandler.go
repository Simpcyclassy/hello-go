package handlers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Simpcyclassy/hello-go/cache"
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
	cacheBody, err := h.memCache.Get(q)
	if err != nil || cacheBody == nil {
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
			log.Error().Err(err).Msg("Problem with unmarshalling")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops something went wrong, please try again later"))
			return
		}

	} else {
		informationBody, err = getBytes(cacheBody)
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

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
