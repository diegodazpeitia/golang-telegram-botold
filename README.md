# Golang Telegram Bot

This is a simple Telegram bot built with Golang using the `telebot.v3` library.

## Features
- Responds to `/start` command with a welcome message.
- Echoes any text message sent to the bot.

## Setup

1. Clone the repository.
2. Set up your environment variable `TELEGRAM_BOT_TOKEN` with your Telegram bot token.

## Running the Bot

```bash
go run main.go
```

Make sure to set the bot token in your environment:

```bash
export TELEGRAM_BOT_TOKEN=your_bot_token
```

---

That's it! Just replace `your_bot_token` with your actual bot token.
