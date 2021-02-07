package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// InformationResponse creates a struct for the weather body
type InformationResponse struct {
	Weather     WeatherResponse
	CovidReport CovidStats
}

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

// CovidStats creates a struct for covid 19 live stats
type CovidStats struct {
	Country        string `json:"Country"`
	ConfirmedCases int32  `json:"Confirmed"`
	Deaths         int32  `json:"Deaths"`
	Recove         int32  `json:"Recovered"`
	Active         int32  `json:"Active"`
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

const weatherBASE = "http://api.openweathermap.org/data/2.5/weather/"
const covidBASE = "https://api.covid19api.com/live/country"

func checkUrls(urls []string) InformationResponse {
	// var wg sync.WaitGroup
	// var requestBody WeatherResponse

	// for _, link := range urls {
	// 	wg.Add(1)
	// 	go makeWeatherRequests(link, ch)
	// }
	weatherChannel := make(chan WeatherResponse)

	go makeWeatherRequests(urls[0], weatherChannel)
	// go func() {
	// 	wg.Wait()
	// 	requestBody = <-ch
	// 	close(ch)
	// }()
	requestBody := <-weatherChannel
	information := InformationResponse{}
	information.Weather = requestBody
	return information

	// Transformation of our result data
	// return our data structure ready to be put in the response

}

func makeWeatherRequests(url string, c chan WeatherResponse) {
	weather := WeatherResponse{}

	resp, err := http.Get(url)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("Error calling weather API: %s", err.Error()))
		c <- weather
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Info().Msg(err.Error())
		c <- weather
		return
	}

	log.Info().Msg(fmt.Sprintf("weather: %s", body))

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Info().Msg(err.Error())
		c <- weather
		return
	}
	weather.CreatedAt = time.Now()
	c <- weather
}

// InformationHandler handles weather and covid requests...
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	live := time.Now().Add(-24 * time.Hour)
	country := strings.ReplaceAll(q, " ", "-")
	weatherURL := fmt.Sprintf("%s?q=%s&appid=%s", weatherBASE, q, goDotEnvVariable("APP_ID"))
	covidURL := fmt.Sprintf("%s/%s/status/confirmed/date/%s", covidBASE, country, live)

	links := []string{
		weatherURL,
		covidURL,
	}

	weather := checkUrls(links)
	// write the results to our ResponseWriter, if theres an error return this.
	// 	w.Write(informationBody)

	weatherBody, err := json.Marshal(weather)
	if err != nil {
		log.Info().Msg(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(weatherBody)
}
