package main

import (
	"AMA_bot/internal/telegramClient"
	"log"
	"os"
)

// возвращает токен при вводе в cmd
func GetToken() string {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("token is required")
	}
	return token
}

func main() {
	telegramClient.TgClient(GetToken())
}
