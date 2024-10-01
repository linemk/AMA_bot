package telegramClient

import "fmt"

// команда для /start
func StartParser(client *Client, Id int, name string) {
	var text string = fmt.Sprintf("Привет, %s!\nЧтобы узнать погоду в городе -> введи название города\n\nПример: Москва", name)
	err := client.SendMessage(Id, text)
	if err != nil {
		_ = fmt.Errorf("can't send message")
	}
}
