package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// нужна для обработки входящего ответа
type WeatherAnswer struct {
	City          string  `json:"city"`
	Temperature   int     `json:"temperature"`
	Precipitation string  `json:"precipitation"`
	Humidity      int     `json:"humidity"`
	Wind          float64 `json:"wind"`
}

// создадим фиктивные данные о погоде
var Weather = WeatherAnswer{
	City:          "Москва",
	Temperature:   25,
	Precipitation: "Ясно",
	Humidity:      60,
	Wind:          5.5,
}

// парсит ответ от WEATHER API. Форматируем данные для отправки пользователю в текстовом виде
func parseWeatherAnswer(weather WeatherAnswer) string {
	// пока создаем фиктивные данные, которые потом получим от WEATHER API для тестов
	return fmt.Sprintf(
		"🏙 Город: %s\n🌡️ Температура: %d°C\n☀ Осадки: %s\n💧 Влажность: %d%%\n💨 Ветер: %.2f м/с",
		weather.City, weather.Temperature, weather.Precipitation, weather.Humidity, weather.Wind,
	)
}

// Непосредственно сама отправка сообщения в бота
func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest("sendMessage", q)

	if err != nil {
		return fmt.Errorf("can't do request: %v", err)
	}
	return nil
}

// Головная функция отправки ответа
func answerForUser(client *Client, chatID int64, weather WeatherAnswer) {
	// Формируем текст сообщения
	message := parseWeatherAnswer(weather)

	// Отправляем сообщение
	err := client.SendMessage(int(chatID), message)
	if err != nil {
		fmt.Printf("Ошибка при отправке сообщения: %v\n", err)
	}
}
