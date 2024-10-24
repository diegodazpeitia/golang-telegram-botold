package main

import (
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

func main() {
	// Get the bot token from the environment variable
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
		return
	}

	// Set up bot settings
	pref := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// Create the bot
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Handle the /start command
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Welcome to my Telegram bot!")
	})

	// Handle any text messages
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		// Echo the received message back to the user
		return c.Send("You said: " + c.Message().Text)
	})

	// Start the bot
	bot.Start()
}
