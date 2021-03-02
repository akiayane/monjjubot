package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"monjjubot/server"
	"os"
)

func main(){

	//load global variables
	_ = godotenv.Load("globals.env")

	//initialize bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error occurred while creating a bot")
		fmt.Println(err)
	}


	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	channel, err := bot.GetUpdatesChan(u)

	if err != nil {
		fmt.Println("Error occurred while channeling")
		fmt.Println(err)
	}
	commands_handler:= Commands_handler_init(bot, &u,&channel)
	log_handler:= Log_handler_init(bot, &u,&channel)

	bot_handler:= Bot_handler_init(commands_handler,log_handler)

	server.ListenAndServe(bot_handler,channel)

}
