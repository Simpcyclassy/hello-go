package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// InformationResponse creates a struct for the weather and covid data for a country
type InformationResponse struct {
	Weather     WeatherResponse
	CovidReport CovidStats
}

// WeatherResponse creates a struct for the weather data
type WeatherResponse struct {
	Coord     Coords          `json:"coord"`
	Weather   []WeatherResult `json:"weather"`
	Base      string          `json:"base"`
	Main      Measurement     `json:"main"`
	Wind      WindResult      `json:"wind"`
	Sys       System          `json:"sys"`
	CreatedAt time.Time
}

// CovidStats creates a struct for covid 19 live stats
type CovidStats struct {
	Country        string `json:"country"`
	ConfirmedCases int32  `json:"cases"`
	TodaysCases    int32  `json:"todayCases"`
	TotalDeaths    int32  `json:"deaths"`
	TodaysDeaths   int32  `json:"todayDeaths"`
	Recovered      int32  `json:"recovered"`
	Active         int32  `json:"active"`
	Critical       int32  `json:"critical"`
	Tests          int32  `json:"tests"`
}

// Coords creates a struct for weather longitude and latitude
type Coords struct {
	Longitude float32 `json:"lon"`
	Lattitude float32 `json:"lat"`
}

// WeatherResult creates a struct for weather condition and description
type WeatherResult struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// WindResult creates a struct for weather spead and degree
type WindResult struct {
	Speed  float32 `json:"speed"`
	Degree int32   `json:"deg"`
}

// Measurement creates a struct for external weather parameter
type Measurement struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int32   `json:"pressure"`
	Humidity  int32   `json:"humidity"`
}

// System creates a struct for weather locations and time of the day
type System struct {
	Type    int32   `json:"type"`
	ID      int32   `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int32   `json:"sunrise"`
	Sunset  int32   `json:"sunset"`
}

func goDotEnvVariable(key string) string {
	mydir, err := os.Getwd()
	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Debug().Msgf("Logs from go Dot Env %s", mydir)

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	return os.Getenv(key)
}

const weatherBASE = "http://api.openweathermap.org/data/2.5/weather/"
const covidBASE = "https://corona.lmao.ninja/v2/countries/"

func checkUrls(weatherURL, covidURL string) (InformationResponse, error) {

	weatherChannel := make(chan []byte)
	covidChannel := make(chan []byte)
	go makeCountryInfoRequests(weatherURL, weatherChannel)
	go makeCountryInfoRequests(covidURL, covidChannel)
	weatherBody := <-weatherChannel
	covidBody := <-covidChannel

	var weatherInfo WeatherResponse
	err := json.Unmarshal(weatherBody, &weatherInfo)
	if err != nil {
		log.Debug().Msg("Error with unmarshelling the Weather Info")
		return InformationResponse{}, err
	}
	weatherInfo.CreatedAt = time.Now()

	var covidInfo CovidStats
	err = json.Unmarshal(covidBody, &covidInfo)
	if err != nil {
		log.Debug().Msg("Error with unmarshelling the Covid Info")
		return InformationResponse{}, err
	}

	var information InformationResponse
	information.Weather = weatherInfo
	information.CovidReport = covidInfo
	return information, nil

}

func makeCountryInfoRequests(url string, c chan []byte) {
	var countryInfo []byte

	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		// we might want to redact sensitive information like the APP ID
		// helper function that replaces the APP ID with ****
		log.Error().Msg(fmt.Sprintf("Error calling API URL %s, status code: %d ", url, resp.StatusCode))
		c <- countryInfo
	}
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error calling API URL %s %s", url, err.Error()))
		c <- countryInfo
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Request Body is empty: %s", err.Error()))
		c <- countryInfo
		return
	}
	resp.Body.Close()
	if err != nil {
		log.Error().Err(err).Msg("CHIOMA")
		c <- countryInfo
		return
	}
	c <- body
}

// InformationHandler handles weather and covid requests...
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	if len(q) < 1 {
		// Add metric for param to see how many people who make reuests without a query
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not passing a country param"))
		return
	}
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
