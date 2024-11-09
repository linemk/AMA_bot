package telegramClient

import (
	translate "AMA_bot/pkg/translateAPI" // Импортируйте пакет с переводом
	weather "AMA_bot/pkg/weatherAPI"     // Импортируйте пакет с погодой
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

// парсит ответ от WEATHER API. Форматируем данные для отправки пользователю в текстовом виде
func parseWeatherAnswer(weather weather.WeatherAnswer) string {
	// Мапа для смайлов осадков
	smilesMap := map[string][]string{
		"🌤":  {"Без осадков", "Переменная облачность"},
		"☁️": {"Пасмурно", "Облачно"},
		"🌧":  {"Небольшой дождь", "Умеренный дождь", "Дождь", "Ливень", "Шторм", "Небольшой ливневый дождь", "Умеренный или сильный ливневый дождь"},
		"❄️": {"Снег", "Метель", "Небольшой снег"},
		"🌫":  {"Туман", "Дымка"},
	}
	var smileIcon string = "☀"

	for i, v := range smilesMap {
		for _, precipitation := range v {
			if precipitation == weather.Precipitation {
				smileIcon = i
			}
		}
	}

	// Убираем лишние символы из города
	resultCity := ""
	cityName := translate.EngToRus(weather.City)
	if strings.Contains(cityName, "q") || strings.Contains(cityName, "langpair") {
		cityName = strings.ReplaceAll(cityName, "q", "")
		cityName = strings.ReplaceAll(cityName, "langpair", "")
		cityName = strings.ReplaceAll(cityName, "en", "")
		cityName = strings.ReplaceAll(cityName, "ru", "")

	}
	if strings.Contains(cityName, "Re") {
		cityName = strings.ReplaceAll(cityName, "Re", "")

	}

	for i, symbol := range cityName {
		if i == 0 && symbol == 'г' {
			continue
		}
		if !unicode.IsLetter(symbol) && symbol != '-' && symbol != ' ' {
			continue
		}
		resultCity += string(symbol)
	}

	result := fmt.Sprintf(
		"🏙 Город: %s\n🌡️ Температура: %d°C\n%v %s\n💧 Влажность: %d%%\n💨 Ветер: %.2f м/с",
		resultCity, weather.Temperature, smileIcon, weather.Precipitation, weather.Humidity, weather.Wind/3.6)
	if weather.City == "" {
		result = "💫 Возможно звезды не так сошлись...\nПопробуйте изменить запрос или написать город латиницей.\nПример -> Krasnodar"
	}
	return result
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
