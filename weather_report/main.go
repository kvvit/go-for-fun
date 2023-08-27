package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	lat, lon, city := getCoordinates()
	fmt.Printf("Current city: %s\n", city)

	wetherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)
	res, err := http.Get(wetherURL)
	if err != nil {
		fmt.Printf("error making http request for wether: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	response := make(map[string]interface{})
	json.Unmarshal(resBody, &response)

	fmt.Printf("Temperature: %.1f\n", response["current_weather"].(map[string]interface{})["temperature"].(float64))
	fmt.Printf("Wind: %.1f\n", response["current_weather"].(map[string]interface{})["windspeed"].(float64))
}

func getCoordinates() (float64, float64, string) {
	serverPort := 80
	requestURL := fmt.Sprintf("http://ip-api.com:%d/json", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request for coordinates: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var response map[string]any
	json.Unmarshal(resBody, &response)
	return response["lat"].(float64), response["lon"].(float64), response["city"].(string)
}
