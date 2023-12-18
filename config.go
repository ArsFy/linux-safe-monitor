package main

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigType struct {
	// Support: en/cn
	Language string `json:"lang"`
	// -1: No remind and kill
	// 0: Only remind
	// 1: Remind and kill (White list)
	// 2: Remind and kill (All)
	KillMode int `json:"kill_mode"`
	// Interval time (ms)
	CheckTime int `json:"check_time"`

	// Push
	// Telegram
	EnableTelegram bool   `json:"enable_telegram"`
	TelegramToken  string `json:"telegram_token"`
	TelegramChatID int64  `json:"telegram_chat_id"`
}

// Default config
var Config = ConfigType{
	Language:       "en",
	KillMode:       0,
	CheckTime:      5000,
	EnableTelegram: false,
}

// White list
var WhiteList = make([]string, 0)

func init() {
	config, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("Config not found, useing default config.")
		configByte, _ := json.Marshal(Config)
		os.WriteFile("config.json", configByte, 0644)
	}
	json.Unmarshal(config, &Config)

	whiteList, err := os.ReadFile("white-list.json")
	if err != nil {
		log.Println("White list not found, using empty white list.")
	}
	json.Unmarshal(whiteList, &WhiteList)

	// Start
	if Config.EnableTelegram {
		go startTelegram()
	}
}
