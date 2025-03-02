package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	type Weather struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`

		Current struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
			} `json:"condition"`

			Humidity float64 `json:"humidity"`
		} `json:"current"`

		// Forecast struct {
		// 	Forecastday []struct {
		// 		Hour []struct {
		// 			TimeEpoch int64   `json:"time_epoch"`
		// 			TempC     float64 `json:"temp_c"`
		// 			Condition struct {
		// 				Text string `json:"text"`
		// 			} `json:"condition"`
		// 			ChanceOfRain float64 `json:"chance_of_rain"`
		// 		} `json:"hour"`
		// 	} `json:"forecastday"`
		// } `json:"forecast"`
	}

	q := "<Home City>"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	// fmt.Println("App Starting..")
	api_key := "5d2ecd4c69d941a1886163621250103"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", api_key, q)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	weather := Weather{}
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}
	// fmt.Println(weather)
	location, current := weather.Location, weather.Current

	// fmt.Printf("Location: %s, %s\n", location.Name, location.Country)
	// fmt.Printf("Current Temp and Humidity: %.0f, %.0f\n", current.TempC, current.Humidity)
	fmt.Printf("%s ,%s, %.0fC, %.0f%%\n", location.Name, location.Country, current.TempC, current.Humidity)

}
