package handlerstest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Simpcyclassy/hello-go/handlers"
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
			path:               "/tree",
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
			message:            "Bad request, please go to /tree",
			paramValue:         "oak",
			expectedErr:        false,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "third test - wrong method",
			method:             "POST",
			path:               "/tree",
			param:              "favoriteTree",
			message:            "404 Method not found.",
			paramValue:         "oak",
			expectedErr:        false,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "third test - empty query",
			method:             "GET",
			path:               "/tree",
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
			handlers.FavoriteTreeHandler(w, req)
			resp := w.Result()

			if resp.StatusCode != testCase.expectedStatusCode {
				t.Errorf("got %v, expected %v", resp.StatusCode, testCase.expectedStatusCode)
			}
		})
	}
}
