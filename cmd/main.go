package main

import (
	"AMA_bot/internal/telegramClient"
	"flag"
	"log"
)

// возвращает токен при вводе в cmd
func GetToken() string {
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

func main() {
	telegramClient.TgClient(GetToken())
}
