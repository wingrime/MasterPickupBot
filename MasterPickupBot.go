package main

import (
	"log"
	"os"
	"time"

	"github.com/tucnak/telebot"
)

const (
	BotApiKey = "API_KEY" // "Telegram API Key"
)

var (
	logger = log.New(os.Stderr, "[main] ", log.LstdFlags)
)

func main() {
	SecretAPIKey := os.Getenv(BotApiKey)
	if SecretAPIKey == "" {
		logger.Fatalf("env variable %s not specifited", BotApiKey)
	}
	bot, err := telebot.NewBot(SecretAPIKey)
	if err != nil {
		logger.Fatal(err)
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/start" {
			bot.SendMessage(message.Chat, "Привет, я раскажу тебе как двигать тазом, "+message.Sender.FirstName+"!", nil)
		}
	}
}
