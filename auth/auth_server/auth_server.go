package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/auth"
	"monjjubot/database"
	"net"
	"os"
)

type Server struct {
	auth.UnimplementedAuthServiceServer
	router []Pattern
}

type Pattern struct {
	command string
	user_type string
}

func (s *Server) AuthService(ctx context.Context,request *auth.AuthRequest) (*auth.AuthResponse,error){
	router := s.router
	response := &auth.AuthResponse{Status: false,Message: "Allowed"}
	chat_id := request.ChatId
	command := request.Command
	for i:=0;i<len(router);i++ {
		if router[i].command==command{
			if router[i].user_type=="all" {
				response.Status=true
				response.Message="Allowed"
			} else {
				status := false
				status, _ = checkVerification(chat_id)
				response.Status = status
			}
		}
	}
	return response, nil
}


func checkVerification(chat_id string) (bool,error){
	_ = godotenv.Load("globals.env")
	port := os.Getenv("DATABASE_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := database.NewDatabaseAccessServiceClient(conn)
	ctx := context.Background()
	req := &database.VerificationRequest{Chatid: chat_id}
	response, err := c.CheckVerification(ctx, req)
	return response.Status, err
}

func main(){
	_ = godotenv.Load("globals.env")
	port := os.Getenv("AUTH_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	l, err := net.Listen("tcp", address+":"+port)
	if err!=nil{
		log.Fatalln("Error occurred while deploying the server",err)
	}
	s := grpc.NewServer()
	router:=[]Pattern{
		{
			command:   "/start",
			user_type: "all",
		},
		{
			command:   "/get",
			user_type: "authorized",
		},
		{
			command:   "/reg",
			user_type: "all",
		},
		{
			command:   "/guide",
			user_type: "all",
		},
	}
	auth.RegisterAuthServiceServer(s, &Server{router: router})
	log.Println("AUTH_Server is running on port: "+port)
	if err := s.Serve(l); err!=nil{
		log.Fatalln("Error occured while listening for auth server", err)
	}
}

