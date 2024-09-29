package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// нужна для обработки входящего ответа
type WeatherAnswer struct {
	City          string  `json:"city"`
	Temperature   int     `json:"temperature"`
	Precipitation string  `json:"precipitation"`
	Humidity      int     `json:"humidity"`
	Wind          float64 `json:"wind"`
}

// структура ответа
type SendMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// функция парсит ответ от WEATHER API. Форматируем данные для отправки пользователю в текстовом виде
func parseWeatherAnswer(weather WeatherAnswer) string {
	return fmt.Sprintf(
		"Город: %s\nТемпература: %d°C\nОсадки: %s\nВлажность: %d%%\nВетер: %.2f м/с",
		weather.City, weather.Temperature, weather.Precipitation, weather.Humidity, weather.Wind,
	)
}

// Основная функция отправки ответа
func main() {
	botToken := "7699031903:AAFRtoi4vh2i12MvuSPreBd-x5KHlVQPf_M" // мне его надо где-то взять подгрузить из .env

	// тело ответа, согласно структуре
	reqBody := SendMessage{
		ChatID: chatID,                      // нужно получить его от Макса (первоначальный запрос)
		Text:   parseWeatherAnswer(weather), // нужно получить weather от Андрея (weatherApi)
	}

	// превращаем уже готовый ответ в последовательность байтов для отправки далее
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return
	}
	// url адрес для ответа ниже в POST запросе
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// сам пост запрос обращение к телеграм апи и передача ему тела запроса
	answer, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}
	// закрытие тела запроса
	defer answer.Body.Close()

	// проверка на ошибку
	if answer.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: %v\n", answer.Status)
	}

}
