package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"google.golang.org/grpc"
	"log"
	"monjjubot/request"
)



func connect(pattern string) (response string){

	fmt.Println("ClientConnected")

	conn, err := grpc.Dial("localhost:59751", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := request.NewTelegramBotServiceClient(conn)
	response = doRequest(c, pattern)
	return response

}

func doRequest(client request.TelegramBotServiceClient, pattern string) (response string){
	ctx := context.Background()
	req := &request.CommandPackRequest{Command:pattern}

	res, err := client.CommandPack(ctx, req)
	if err!=nil{
		log.Fatalln("Error when getting response from server: ",err)
	}
	response = res.Response
	return response
}


func Bot_handler_init(bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln("Error when listening to updates: ",err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		response_string := connect(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response_string)
		_, err = bot.Send(msg)
		if err!=nil{
			log.Println("Error when sending via tg bot: ",err)
		}

	}

}






