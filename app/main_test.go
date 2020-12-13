package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestConverter(t *testing.T) {

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
			path:               "/",
			param:              "favoriteTree",
			message:            "It's nice to know that your favorite tree is a oak",
			paramValue:         "oak",
			expectedErr:        false,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "second test - wrong path",
			method:             "GET",
			path:               "/wrong",
			param:              "favoriteTree",
			message:            "Bad request, please go to /",
			paramValue:         "oak",
			expectedErr:        false,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "third test - wrong method",
			method:             "POST",
			path:               "/",
			param:              "favoriteTree",
			message:            "404 Method not found.",
			paramValue:         "oak",
			expectedErr:        false,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "third test - empty query",
			method:             "GET",
			path:               "/",
			param:              "favoriteTree",
			message:            "Please tell me your favorite tree",
			paramValue:         "",
			expectedErr:        false,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, _ := http.NewRequest(testCase.method, testCase.path, nil)

			query := req.URL.Query()
			query.Add(testCase.param, testCase.paramValue)
			req.URL.RawQuery = query.Encode() // Research this

			w := httptest.NewRecorder()
			favoriteTreeHandler(w, req)
			resp := w.Result()

			body, _ := ioutil.ReadAll(resp.Body)
			if resp.StatusCode != testCase.expectedStatusCode {
				t.Errorf("got %v, expected %v", resp.StatusCode, testCase.expectedStatusCode)
			}
			fmt.Println("tesCase:", reflect.TypeOf(testCase.message))
			fmt.Println("body:", reflect.TypeOf(body))

			if string(body) == testCase.message {
				t.Errorf("got %v, expected %v", string(body), testCase.message)

			}
		})
	}
}
