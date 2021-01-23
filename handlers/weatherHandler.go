package weatherHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type WeatherResponse struct {
	Coord   	Coords          `json:"coord"`
	Weather 	[]WeatherResult `json:"weather"`
	Base    	string          `json:"base"`
	Main		Measurement		`json:"main"`
	Wind		WindResult		`json:"wind"`
	Sys			System			`json:"sys"`
	CreatedAt	time.Time
}
type Coords struct {
	Longitude float32 `json:"lon"`
	Lattitude float32 `json:"lat"`
}

type WeatherResult struct {
	ID          int         `json:"id"`
	Main        string		`json:"main"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
}

type WindResult struct {
	Speed         float32      `json:"speed"`
	Degree        int32		   `json:"deg"`
}

type Measurement struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int32   `json:"pressure"`
	Humidity  int32   `json:"humidity"`
}

type System struct {
	Type      int32 	`json:"type"`
	ID 		  int32 	`json:"id"`
	Message   float32 	`json:"message"`
	Country   string 	`json:"country"`
	Sunrise   int32   	`json:"sunrise"`
	Sunset    int32   	`json:"sunset"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}
const weatherURL = "http://api.openweathermap.org/data/2.5/weather"

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	resp, err := http.Get(fmt.Sprintf("%s?q=%s&appid=%s", weatherURL, q, goDotEnvVariable("APP_ID")))

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)

	weather := WeatherResponse{}
	weather.CreatedAt = time.Now()
	err = json.Unmarshal(body, &weather)
	fmt.Printf("%v", weather)

	weatherBody, err := json.Marshal(weather)
	if err != nil {
		fmt.Printf("weatherError :", err)
	}
	w.Write(weatherBody)
}
