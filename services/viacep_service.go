package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ViaCepResponse struct {
	Localidade string `json:"localidade"`
}

func GetLocationByZipCode(zipcode string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch location")
	}

	var data ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", errors.New("location not found")
	}

	return data.Localidade, nil
}
