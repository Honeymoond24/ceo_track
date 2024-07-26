package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type INotifier interface {
	SendMessage(message string) error
}

type Notifier struct {
	ChatID   string
	BotToken string
}

func (n Notifier) SendMessage(message string) error {
	// Send message
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.BotToken)
	data := map[string]interface{}{
		"text":       message,
		"chat_id":    n.ChatID,
		"parse_mode": "HTML",
	}
	jsonData, err := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Message sent", resp)
	return nil
}
