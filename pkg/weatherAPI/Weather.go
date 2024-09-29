package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph  float64 `json:"wind_kph"`
		Humidity int     `json:"humidity"`
	} `json:"current"`
}

// s
type WeatherAnswer struct {
	City          string  `json:"city"`
	Temperature   int     `json:"temperature"`
	Precipitation string  `json:"precipitation"`
	Humidity      int     `json:"humidity"`
	Wind          float64 `json:"wind"`
}

// GetWeather возвращает данные о погоде для указанного города
func GetWeather(city string) WeatherAnswer {
	apiKey := "840b43c108d3402292d160550242909"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherAnswer{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherAnswer{}
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return WeatherAnswer{}
	}

	// Преобразование данных в WeatherAnswer
	answer := WeatherAnswer{
		City:          weather.Location.Name,
		Temperature:   int(weather.Current.TempC),
		Precipitation: weather.Current.Condition.Text,
		Humidity:      weather.Current.Humidity,
		Wind:          weather.Current.WindKph,
	}

	return answer
}
