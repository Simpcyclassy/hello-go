// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type WeatherResults struct {
// 	Coord   Coords          `json:"coord"`
// 	Weather []WeatherResult `json:"weather"`
// 	Base    string          `json:"base"`
// }

// type Coords struct {
// 	Longitude float32 `json:"lon"`
// 	Lattitude float32 `json:"lat"`
// }

// type WeatherResult struct {
// 	ID          int         `json:"id"`
// 	Main        Measurement `json:"main"`
// 	Description string      `json:"description"`
// 	Icon        string      `json:"icon"`
// }

// type Measurement struct {
// 	Temp      float32 `json:"temp"`
// 	FeelsLike float32 `json:"feels_like"`
// 	TempMin   float32 `json:"temp_min"`
// 	TempMax   float32 `json:"temp_max"`
// 	Pressure  int32   `json:"pressure"`
// 	Humidity  int32   `json:"humidity"`
// }

// const appID = "0f36a09d6521ef6fbfe48490de11f0d9"

// func weatherHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/weather" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Bad request, please go to /weather"))
// 		return
// 	}
// 	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=London&appid=%s", appID))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s", body)

// 	weather := WeatherResults{}
// 	err = json.Unmarshal(body, &weather)
// 	fmt.Printf("%v", weather)
// }

// func main() {

// 	http.HandleFunc("/weather", weatherHandler)

// 	log.Println("Starting server at port 8000")
// 	if err := http.ListenAndServe(":8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

type WeatherResults struct {
	Coord   Coords          `json:"coord"`
	Weather []WeatherResult `json:"weather"`
	Base    string          `json:"base"`
}

type Coords struct {
	Longitude float32 `json:"lon"`
	Lattitude float32 `json:"lat"`
}

type WeatherResult struct {
	ID          int         `json:"id"`
	Main        Measurement `json:"main"`
	Description string      `json:"description"`
	Icon        string      `json:"icon"`
}

type Measurement struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int32   `json:"pressure"`
	Humidity  int32   `json:"humidity"`
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
  

func main() {
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=London&appid=%s", goDotEnvVariable("APP_ID")))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)

	weather := WeatherResults{}
	err = json.Unmarshal(body, &weather)
	fmt.Printf("%v", weather)
}
