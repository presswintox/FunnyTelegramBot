package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("810113301:AAGtsQ18hEE9mVcYVNX1uqb7AHzn0RKnwwQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.Text == "Да" || update.Message.Text == "да"{
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пизда")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
		}
	}
}