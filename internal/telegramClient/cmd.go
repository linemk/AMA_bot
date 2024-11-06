package telegramClient

import "fmt"

// команда для /start
func StartParser(client *Client, Id int, name string) error {
	if Id == 0 || name == "" {
		return fmt.Errorf("id or name is empty")
	}
	var text string = fmt.Sprintf("Привет, %s!\nЧтобы узнать погоду в городе -> введи название города\n\nПример: Москва", name)
	err := client.SendMessage(Id, text)
	if err != nil {
		return fmt.Errorf("can't send message")
	}
	return nil
}
