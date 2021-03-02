package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"os"
)

type app struct {
	handler Bot_handler
	channel *tgbotapi.UpdatesChannel
}

func(a *app) get_handler() *Bot_handler {
	return &a.handler
}

func(a *app) bot_app_init() app {

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

	return app{bot_handler, &channel}
}

func(a *app) ListenAndServe(){


	fmt.Println("Bot is running...")
	for update := range *a.channel {
		if update.Message == nil {
			continue
		}

		command, param:= a.handler.separate_param(update.Message.Text)
		a.handler.call_function(command, param, update)

	}
	/*
		fmt.Println(handler.command_pool)*/
}