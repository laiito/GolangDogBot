package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var animalsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("\U0001F436"),
	),
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		file := tgbotapi.FileURL("img2.joyreactor.cc/pics/post/гифки-живность-котэ-рыба-7158036.gif")
		gifMsg := tgbotapi.NewVideo(update.Message.Chat.ID, file)
		gifMsg.ReplyMarkup = animalsKeyboard
		bot.Send(gifMsg)
	}

}
