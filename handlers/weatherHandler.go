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

// WeatherResponse creates a struct for the weather body
type WeatherResponse struct {
	Coord     Coords          `json:"coord"`
	Weather   []WeatherResult `json:"weather"`
	Base      string          `json:"base"`
	Main      Measurement     `json:"main"`
	Wind      WindResult      `json:"wind"`
	Sys       System          `json:"sys"`
	CreatedAt time.Time
}

// Coords creates a struct for weather longitude and latitude
type Coords struct {
	Longitude float32 `json:"lon"`
	Lattitude float32 `json:"lat"`
}

// WeatherResult creates a struct for weather description
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
		fmt.Println(err)
	}
	fmt.Println(mydir)

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	return os.Getenv(key)
}

const weatherURL = "http://api.openweathermap.org/data/2.5/weather"

// WeatherHandler handles favorite weather requests...
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	resp, err := http.Get(fmt.Sprintf("%s?q=%s&appid=%s", weatherURL, q, goDotEnvVariable("APP_ID")))

	if err != nil {
		log.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops something went wrong! Please try again :D"))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	log.Info().Msg(fmt.Sprintf("body: %s", body))

	weather := WeatherResponse{}
	weather.CreatedAt = time.Now()
	err = json.Unmarshal(body, &weather)

	weatherBody, err := json.Marshal(weather)
	if err != nil {
		log.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	w.Write(weatherBody)
}
