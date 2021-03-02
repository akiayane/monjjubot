package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/request"
	"net"
	"os"
	"strings"
)


type Server struct {
	request.UnimplementedTelegramBotServiceServer
}

func separate_param(str string) (string,string) {
	if strings.Contains(str, " "){
		space_index := strings.Index(str," ")
		command := str[0:space_index]
		param := str[space_index:]
		return strings.Trim(command," "), strings.Trim(param," ")

	}
	return strings.Trim(str," "), ""
}

//Greet is an example of unary rpc call
func (s *Server) CommandPack(ctx context.Context,req *request.CommandPackRequest) (*request.CommandPackResponse, error) {
	fmt.Println("Start function invoked")
	command:=req.Command

	res := &request.CommandPackResponse{Status: false, Response: "Unknown command, use /main to see available commands"}

	if command=="/start" {
		res.Status = true
		res.Response = "Hello new user"
	}

	return res, nil
}

func main() {
	_ = godotenv.Load("globals.env")
	port := os.Getenv("MAIN_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	l, err := net.Listen("tcp", address+port)
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	request.RegisterTelegramBotServiceServer(s, &Server{})
	log.Println("Command handler Server is running on port: "+port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}