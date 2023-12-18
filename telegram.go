package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TelegramBot *tgbotapi.BotAPI

func startTelegram() {
	var err error

	TelegramBot, err = tgbotapi.NewBotAPI(Config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}
	TelegramBot.Debug = false
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// updates := bot.GetUpdatesChan(u)
	// for update := range updates {
	// 	if update.Message != nil {

	// 	}
	// }
}
