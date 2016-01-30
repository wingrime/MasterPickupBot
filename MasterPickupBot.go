package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

const (
	BotApiKey = "API_KEY" // "Telegram API Key"
)

var (
	logger = log.New(os.Stderr, "[main] ", log.LstdFlags)
)

func processGeneralMessage(str string, dict []string) string {
	return "Он сказал: " + dict[rand.Intn(len(dict))]
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	dictData, err := ioutil.ReadFile("dict.txt")
	if err != nil {
		logger.Fatal(err)
	}
	dict := strings.Split(string(dictData), "\n")

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
	sendOptions := telebot.SendOptions{
		ReplyMarkup: telebot.ReplyMarkup{
			ForceReply: true,
			Selective:  true,

			CustomKeyboard: [][]string{
				[]string{"Совет Гуру"},
			},
		},
	}
	for message := range messages {
		logger.Printf("Req %s from user %s ", message.Text, message.Sender.FirstName)
		if message.Text == "/start" {
			bot.SendMessage(message.Chat, "Привет, я раскажу тебе как двигать тазом, "+message.Sender.FirstName+"!", &sendOptions)
		} else {
			bot.SendMessage(message.Chat, processGeneralMessage(message.Text, dict), &sendOptions)
		}
	}
}
