package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){
	//load global variables
	_ = godotenv.Load("globals.env")
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil{
		log.Println("Error when creating a bot instance occured", err)
	}
	log.Println("TG bot is running")
	Bot_handler_init(bot)
}