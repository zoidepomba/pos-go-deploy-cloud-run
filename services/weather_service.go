package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByLocation(location string) (float64, error) {
	apiKey := "c419db08b0a8421583e223702240412"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, location)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Weather API Response:", string(body))
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Response Body: %s\n", string(body))
		fmt.Printf("Request URL: %s\n", url)
		return 0, errors.New("failed to fetch weather data")
	}

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
