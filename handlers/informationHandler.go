package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Simpcyclassy/hello-go/cache"
	"github.com/rs/zerolog/log"
)

type informationHandler struct {
	memCache cache.Cache
}

// InformationHandler handles weather and covid requests...
func (h *informationHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	if len(q) < 1 {
		// Add metric for param to see how many people who make reuests without a query
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not passing a country param"))
		return
	}
	_, err := h.memCache.Get(q)

	//check data exists in the cache
	//if truemove to line w.WriteHeader(http.StatusOK)
	//if not the checUrls and marshall information

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

	informationBody, err := json.Marshal(info)
	if err != nil {
		// Internal record of what happened
		log.Error().Err(err).Msg("Problem with unmarshalling")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops something went wrong, please try again later"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(informationBody)
}
