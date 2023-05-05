package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func convert(amount float64) float64 {
	// найти API для ковертации валюты
	return amount * 11.06
}

// Получение сообщения от пользователя
func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg string) {
	if update.Message == nil {
		return
	}

	res := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	bot.Send(res)
}

// обработка ввода пользователя и конвертация валюты
func handleConvert(bot *tgbotapi.BotAPI, update tgbotapi.Update, updates tgbotapi.UpdatesChannel) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, MsgConv)
	bot.Send(msg)

	response := <-updates

	amount, err := strconv.ParseFloat(response.Message.Text, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, MsgNumErr)
		bot.Send(msg)
		return
	}

	res := convert(amount)
	msgres := fmt.Sprintf(MsgRes, res)
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msgres))
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
