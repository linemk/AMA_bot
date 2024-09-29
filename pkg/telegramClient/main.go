package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// возвращает токен при вводе в cmd
func getToken() string {
	token := flag.String(
		"token",
		"",
		"give token from telegram",
	)
	// обрабатывает наше значение
	flag.Parse()
	if *token == "" {
		log.Fatal("token is required")
	}
	return *token
}

// хост является константой
const Host string = "api.telegram.org"

func main() {
	// при merge - удалим
	var token string = "7699031903:AAFRtoi4vh2i12MvuSPreBd-x5KHlVQPf_M"

	// создаем клиента для обработки ответа от тг и отправки структуры на сервер
	client := NewClient(Host, token)

	var offset int // смещение
	for {
		// gолучаем обновления (сообщения) от пользователя
		updates, err := client.Updates(offset, 10)
		if err != nil {
			log.Printf("error while getting updates: %v\n", err)
			time.Sleep(1 * time.Second) // Ждём секунду перед повторной попыткой в случае ошибки
			continue
		}

		// Обрабатываем каждое сообщение в цикле
		for _, update := range updates {
			// Выводим полученное сообщение в консоль (для отладки)
			fmt.Printf("New message from update %d: %v\n", update.Id, update.Message)

			// Отправляем ответ пользователю
			err := client.SendMessage(update.Message.Chat.Id, "Погода в Москве +25")
			if err != nil {
				log.Printf("error while sending message: %v\n", err)
			}
			// Обновляем offset, чтобы не получать старые сообщения повторно
			offset = update.Id + 1
		}

		// Делаем паузу перед следующим запросом обновлений
		time.Sleep(2 * time.Second)
	}
}
