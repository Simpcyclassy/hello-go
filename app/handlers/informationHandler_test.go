package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
)

const (
	mockWeatherResponse = "{\"coord\":{\"lon\":-0.1257,\"lat\":51.5085},\"weather\":[{\"id\":800,\"main\":\"Clear\",\"description\":\"clear sky\",\"icon\":\"01n\"}],\"base\":\"stations\",\"main\":{\"temp\":274.61,\"feels_like\":271.82,\"temp_min\":273.15,\"temp_max\":276.15,\"pressure\":1025,\"humidity\":70},\"visibility\":10000,\"wind\":{\"speed\":0.51,\"deg\":0},\"clouds\":{\"all\":0},\"dt\":1615154254,\"sys\":{\"type\":1,\"id\":1414,\"country\":\"GB\",\"sunrise\":1615098752,\"sunset\":1615139446},\"timezone\":0,\"id\":2643743,\"name\":\"London\",\"cod\":200}"
	mockCovidResponse   = "{\"updated\":1615154357539,\"country\":\"Nigeria\",\"countryInfo\":{\"_id\":566,\"iso2\":\"NG\",\"iso3\":\"NGA\",\"lat\":10,\"long\":8,\"flag\":\"https://disease.sh/assets/img/flags/ng.png\"},\"cases\":158237,\"todayCases\":0,\"deaths\":1964,\"todayDeaths\":0,\"recovered\":137645,\"todayRecovered\":0,\"active\":18628,\"critical\":10,\"casesPerOneMillion\":755,\"deathsPerOneMillion\":9,\"tests\":1580442,\"testsPerOneMillion\":7541,\"population\":209574669,\"continent\":\"Africa\",\"oneCasePerPeople\":1324,\"oneDeathPerPeople\":106708,\"oneTestPerPeople\":133,\"activePerOneMillion\":88.88,\"recoveredPerOneMillion\":656.78,\"criticalPerOneMillion\":0.05}"
	weatherBaseURL      = "http://api.openweathermap.org/data/2.5/weather/?q=%s&appid=%s"
	covidBaseURL        = "https://corona.lmao.ninja/v2/countries/%s"
)

// passing none existing country
// no parameter is passed
// 200 from upstream API but no body -> panic
// 500 response from the upstream API
// incorrect data or malformed data

func TestWeather(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testCases := []struct {
		name                      string
		method                    string
		path                      string
		param                     string
		paramValue                string
		weatherResponseBody       string
		weatherResponseStatusCode int
		covidResponseBody         string
		covidResponseStatusCode   int
		expectedStatusCode        int
	}{
		{
			name:                      "happy path",
			method:                    "GET",
			path:                      "/weather",
			param:                     "q",
			paramValue:                "Nigeria",
			weatherResponseBody:       mockWeatherResponse,
			weatherResponseStatusCode: 200,
			covidResponseBody:         mockCovidResponse,
			covidResponseStatusCode:   200,
			expectedStatusCode:        http.StatusOK,
		},
		{
			name:                      "sad path - covid response has empty body",
			method:                    "GET",
			path:                      "/weather",
			param:                     "q",
			paramValue:                "Nigeria",
			weatherResponseBody:       mockWeatherResponse,
			weatherResponseStatusCode: 200,
			covidResponseBody:         "",
			covidResponseStatusCode:   200,
			expectedStatusCode:        http.StatusInternalServerError,
		},
		{
			name:                      "sad path - weather response has empty body",
			method:                    "GET",
			path:                      "/weather",
			param:                     "q",
			paramValue:                "Nigeria",
			weatherResponseBody:       "",
			weatherResponseStatusCode: 200,
			covidResponseBody:         mockCovidResponse,
			covidResponseStatusCode:   200,
			expectedStatusCode:        http.StatusInternalServerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			weatherURL := fmt.Sprintf(weatherBaseURL, testCase.paramValue, goDotEnvVariable("APP_ID"))
			httpmock.RegisterResponder("GET", weatherURL,
				httpmock.NewStringResponder(testCase.weatherResponseStatusCode, testCase.weatherResponseBody))

			covidURL := fmt.Sprintf(covidBaseURL, testCase.paramValue)
			httpmock.RegisterResponder("GET", covidURL,
				httpmock.NewStringResponder(testCase.covidResponseStatusCode, testCase.covidResponseBody))

			req, _ := http.NewRequest(testCase.method, testCase.path, nil)

			query := req.URL.Query()
			query.Add(testCase.param, testCase.paramValue)
			req.URL.RawQuery = query.Encode()

			w := httptest.NewRecorder()
			// InformationHandler(w, req)
			resp := w.Result()

			if resp.StatusCode != testCase.expectedStatusCode {
				t.Errorf("got %v, expected %v", resp.StatusCode, testCase.expectedStatusCode)
			}
		})
	}
}
