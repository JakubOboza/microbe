package microbe

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	app := NewMicrobe()
	app.initialize()

	router := app.router

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}\n", w.Body.String())
}

func TestEncode64Route(t *testing.T) {

	app := NewMicrobe()
	app.initialize()

	router := app.router

	tests := []struct {
		url, body, response string
	}{
		{url: "/api/base64/encode", body: "{\"data\":\"something\"}", response: "{\"data\":\"c29tZXRoaW5n\"}\n"},
		{url: "/api/base64/decode", body: "{\"data\":\"c29tZXRoaW5n\"}", response: "{\"data\":\"something\"}\n"},
	}

	for _, sample := range tests {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", sample.url, strings.NewReader(sample.body))
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, sample.response, w.Body.String())

	}

}
