package main

import (
	weather "AMA_bot/pkg/weatherAPI" // Импортируйте пакет с погодой
	"fmt"
	"net/url"
	"strconv"
)

// парсит ответ от WEATHER API. Форматируем данные для отправки пользователю в текстовом виде
func parseWeatherAnswer(weather weather.WeatherAnswer) string { // Убедитесь, что здесь используется weather.WeatherAnswer
	return fmt.Sprintf(
		"🏙 Город: %s\n🌡️ Температура: %d°C\n☀ Осадки: %s\n💧 Влажность: %d%%\n💨 Ветер: %.2f м/с",
		weather.City, weather.Temperature, weather.Precipitation, weather.Humidity, weather.Wind/3.6) // fixed
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
func answerForUser(client *Client, chatID int64, weatherData weather.WeatherAnswer) { // Используйте weather.WeatherAnswer здесь
	// Формируем текст сообщения
	message := parseWeatherAnswer(weatherData)

	// Отправляем сообщение
	err := client.SendMessage(int(chatID), message)
	if err != nil {
		fmt.Printf("Ошибка при отправке сообщения: %v\n", err)
	}
}
