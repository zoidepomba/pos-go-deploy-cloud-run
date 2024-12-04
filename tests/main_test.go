package tests

import (
	"net/http"
	"net/http/httptest"
	"project/handlers"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetWeatherInvalidZipCode(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather/123", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/weather/{zipcode}", handlers.GetWeather)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("expected status code 422, got %v", status)
	}
}

func TestGetWeatherNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather/00000000", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/weather/{zipcode}", handlers.GetWeather)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("expected status code 404, got %v", status)
	}
}
