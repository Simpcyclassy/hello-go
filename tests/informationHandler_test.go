package handlerstest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Simpcyclassy/hello-go/handlers"
)

func TestWeather(t *testing.T) {

	testCases := []struct {
		name               string
		method             string
		path               string
		param              string
		message            string
		expectedErr        bool
		expectedStatusCode int
		paramValue         string
	}{
		{
			name:               "happy path",
			method:             "GET",
			path:               "/weather",
			param:              "q",
			paramValue:         "London",
			expectedErr:        false,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, _ := http.NewRequest(testCase.method, testCase.path, nil)

			query := req.URL.Query()
			query.Add(testCase.param, testCase.paramValue)
			req.URL.RawQuery = query.Encode()

			w := httptest.NewRecorder()
			handlers.InformationHandler(w, req)
			resp := w.Result()

			if resp.StatusCode != testCase.expectedStatusCode {
				t.Errorf("got %v, expected %v", resp.StatusCode, testCase.expectedStatusCode)
			}
			fmt.Println("resp.StatusCode", resp.StatusCode)
			fmt.Println("testCase.expectedStatusCode", testCase.expectedStatusCode)

		})
	}
}
