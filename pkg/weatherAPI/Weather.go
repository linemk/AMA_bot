package weather

import (
	translate "AMA_bot/pkg/translateAPI"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
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

// checkTrueCity для проверки города, который отправляем в погодный АПИ
func checkTrueCity(city string) string {
	result := ""

	if utf8.RuneCountInString(city) <= 2 { // Проверка на длину строки 3 и более символов
		return result
	}

	var cityForTranslate string
	for i, symbol := range city {
		if !unicode.IsLetter(symbol) && symbol != '-' && symbol != ' ' { // Проверка символов на буквы
			return result
		}
		if i == 0 {
			cityForTranslate += string(unicode.ToUpper(symbol)) // Первая буква заглавная
		} else {
			cityForTranslate += string(unicode.ToLower(symbol)) // Остальные буквы маленькие
		}
	}

	result = cityForTranslate

	if slices.Contains(Garbage, strings.ToLower(result)) {
		result = ""
	}
	return result
}

// GetWeather возвращает данные о погоде для указанного города
func GetWeather(city string) WeatherAnswer {
	apiKey := "840b43c108d3402292d160550242909" // ключ от погодного api
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=ru", apiKey, translate.RuToEng(checkTrueCity(city)))

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
