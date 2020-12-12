# Telegram document sender

Simple tool, which helps you to send some files from server directly to your telegram

## Installation

Clone Repository, run
```bash
go build main.go -o tgsender
```

## Usage

Run file with arguments:
- *filename* argument. relative file from running directory
- *chatId* argument. Chat where to send file. Can be set via .env file (TELEGRAM_CHAT_ID option)
- *botToken* argument. You telegram bot token, received via [@BotFather](https://t.me/botfather). Can be set via .env file (TELEGRAM_BOT_TOKEN option)

```bash
# Usage with .env file with bot token and chat it set by default
./tgsender filename 

# Usage with .env file with only bot token set by default
./tgsender filename <CHAT_ID>

# Usage without .env file
./tgsender filename <CHAT_ID> <BOT_TOKEN>
```

## Contribution

Feel free to contact with - https://amorev.ru/contact
My Telegram Channel - [@amorev94](https://t.me/amorev94)

## Need help

I need help with creating github releases for all OS for more easy setup guide