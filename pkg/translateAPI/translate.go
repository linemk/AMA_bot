package translateAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура для хранения ответа от переводчика
type TranslationResponse struct {
	ResponseCity ResponseData `json:"responseData"`
}

// Структура для хранения переведенного текста
type ResponseData struct {
	TranslatedText string `json:"translatedText"`
}

// Из русского в английский
func RuToEng(textCity string) string {
	query := fmt.Sprintf("?q=%s&langpair=ru|en&mt=1", textCity) // Формируем строку запроса к этому говнопереводчику
	return serverPartTranslate(query)
}

// Из английского в русский
func EngToRus(textCity string) string {
	query := fmt.Sprintf("?q=%s&langpair=en|ru", textCity) // Формируем строку запроса к этому говнопереводчику
	return serverPartTranslate(query)
}

/*
серверная часть функции, пробывал кучу разных, но этот API хотя бы работает хоть как-то
как вариант попробывать еще сделать проверку, если не нашел город по переводчику, то сделать транскрипцию, это либо новый апи
либо самим как-то извратиться.
*/
func serverPartTranslate(query string) string {
	url := "https://api.mymemory.translated.net/get"
	// Отправляем GET-запрос
	resp, err := http.Get(url + query)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return query
	}
	defer resp.Body.Close() // закрываем тело

	// Читаем ответ от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return query
	}

	// Десериализуем JSON-ответ
	var doneTranslate TranslationResponse
	if err := json.Unmarshal(body, &doneTranslate); err != nil {
		fmt.Println("Ошибка", err)
		return query
	}
	return doneTranslate.ResponseCity.TranslatedText
}
