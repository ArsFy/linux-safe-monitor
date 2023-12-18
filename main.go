package main

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	go checkProcesses()

	fmt.Println("Monitor is running.")
	select {}
}

func checkProcesses() {
	if Config.KillMode != -1 {
		err := getAllProcesses()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		time.AfterFunc(time.Duration(Config.CheckTime)*time.Millisecond, checkProcesses)
	}
}

func push(state int, name, cmdline, pid, user string) {
	var msg string
	switch state {
	case 0:
		msg = fmt.Sprintf(i18n[Config.Language]["process_killed"], name, cmdline, user, pid)
	case 1:
		msg = fmt.Sprintf(i18n[Config.Language]["process_running"], name, cmdline, user, pid)
	}

	if Config.EnableTelegram {
		TelegramBot.Send(tgbotapi.NewMessage(Config.TelegramChatID, msg))
	}
}
