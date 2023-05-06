package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	})

	if err != nil {
		log.Printf("Ошибка получения апдейтов: %v", err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start":
			handleMessage(bot, update, MsgHello)
		case "convert":
			handleConvert(bot, update, updates)
		case "contact":
			handleMessage(bot, update, MsgContact)
		case "order":
			handleMessage(bot, update, MsgOrder)
		default:
			handleMessage(bot, update, MsgErr)
		}
	}
}
