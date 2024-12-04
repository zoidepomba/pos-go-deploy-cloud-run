package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/services"
	"project/utils"
	"strings"

	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcode := strings.TrimSpace(mux.Vars(r)["zipcode"])

	// Validação do CEP
	if len(zipcode) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Busca a localização pelo CEP
	location, err := services.GetLocationByZipCode(zipcode)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		log.Println(location)
		return
	}

	// Busca o clima pela localização
	tempC, err := services.GetTemperatureByLocation(zipcode)
	if err != nil {
		http.Error(w, "failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	// Converte as temperaturas
	tempF := utils.CelsiusToFahrenheit(tempC)
	tempK := utils.CelsiusToKelvin(tempC)

	// Resposta de sucesso
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	})
}
