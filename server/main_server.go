package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/databaseproto"
	"monjjubot/mailer"
	"monjjubot/request"
	"net"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	request.UnimplementedTelegramBotServiceServer

	server_address string
	server_domain_name string
	website_server_port string
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

func connectToMailer(email string, message string) (response bool){

	fmt.Println("ClientConnectedToMailer")
	_ = godotenv.Load("globals.env")
	port := os.Getenv("MAILER_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := mailer.NewEmailSendingServiceClient(conn)
	ctx := context.Background()
	req := &mailer.EmailRequest{Email: email, Message: message}

	res, err := c.SendEmail(ctx, req)
	if err!=nil{
		log.Fatalln("Error when getting response from mailer_server: ",err)
	}
	response = res.Status
	return response

}

func connectToDatabase(course_id int64, subject string) *databaseproto.BigResponse {
	fmt.Println("ClientConnectedToDatabase")
	_ = godotenv.Load("globals.env")
	port := os.Getenv("DATABASE_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c:=databaseproto.NewDatabaseAccessServiceClient(conn)
	ctx := context.Background()
	req := &databaseproto.BookRequest{CourseId: course_id, Subject: subject}
	response, _ := c.CourseRequest(ctx,req)
	return response
}


func (s *Server) CommandPack(ctx context.Context,req *request.CommandPackRequest) (*request.CommandPackResponse, error) {
	command:=req.Command
	param := ""
	command,param = separate_param(command)

	res := &request.CommandPackResponse{Status: false, Response: "Unknown command, use /guide to see available commands"}

	switch command {
		case "/get":
			if param==""{
				res.Status = true
				res.Response = "the Course parameter is empty"
				break
			}
			res.Status = true
			course_id, _ := strconv.ParseInt(param, 10, 32)
			response_books := connectToDatabase(course_id, "")
			books_array := response_books.BookPacks
			res.Response = ""
			for _,s := range books_array{
				res.Response = res.Response + s.Subject + s.BookLink + s.BookLink + "\n"
			}

		case "/start":
			res.Status = true
			res.Response = "Hello new user"
		case "/guide":
			res.Status = true
			res.Response = "Available commands: /guide, /start, /reg <example_email@astaneit.edu.kz>"
		case "/reg":
			if param==""{
				res.Status = true
				res.Response = "Email field is empty, use /guide to see available commands"
				break
			}
			if strings.Contains(param, "@astanait.edu.kz"){
				res.Status = true
				email_address := param
				vkey := md5.Sum([]byte(email_address))
				log.Println(vkey)
				fmt.Printf("   MD5: %x\n", vkey)
				//vkey := "123124"
				email_message := "Hello, dear user! " +
					"\nYou see this message because you have registered your email in Monjjubot. " +
					"\nTo finish registration please follow this link: \n" +
					s.server_domain_name+":"+s.website_server_port+"/verify/"+hex.EncodeToString(vkey[:])+"\n" +
					"\n If you did not do that, please close this letter and do not respond. "
				sending_status := connectToMailer(email_address, email_message)
				if sending_status{
					res.Response = "Email accepted"
				}else{
					res.Response = "Error occured while sending email, pls connect to devs"
				}

			}else{
				res.Status = true
				res.Response = "Email is invalid, or do not belongs to @astanait.edu.kz domen"
			}
	}

	return res, nil
}

func main() {
	_ = godotenv.Load("globals.env")
	port := os.Getenv("MAIN_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	l, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	request.RegisterTelegramBotServiceServer(s, &Server{
		server_address: address,
		website_server_port: os.Getenv("WEBSITE_PORT"),
		server_domain_name: os.Getenv("SERVER_DOMAIN_NAME"),
	})
	log.Println("Command handler Server is running on port: "+port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}