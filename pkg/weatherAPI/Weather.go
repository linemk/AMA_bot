package weather

import (
	translate "AMA_bot/pkg/translateAPI"
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

// структура ответа
type WeatherAnswer struct {
	City          string  `json:"city"`
	Temperature   int     `json:"temperature"`
	Precipitation string  `json:"precipitation"`
	Humidity      int     `json:"humidity"`
	Wind          float64 `json:"wind"`
}

// GetWeather возвращает данные о погоде для указанного города
func GetWeather(city string) WeatherAnswer {
	apiKey := "840b43c108d3402292d160550242909" // ключ от погодного api
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=ru", apiKey, translate.RuToEng(city))

	resp, err := http.Get(url) // создаём запрос к api
	if err != nil {
		return WeatherAnswer{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // получаем ответ
	if err != nil {
		return WeatherAnswer{}
	}

	var weather WeatherResponse // парсим это в структуру
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
