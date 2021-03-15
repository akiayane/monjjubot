package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/database"
	"monjjubot/mailer"
	"monjjubot/main_server"
	"monjjubot/auth"
	"net"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	main_server.UnimplementedTelegramBotServiceServer

	server_address string
	server_domain_name string
	website_server_port string
}

func Auth_service(chatid string,command string) (string,bool) {
	_ = godotenv.Load("globals.env")
	port := os.Getenv("AUTH_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	ctx := context.Background()
	req := &auth.AuthRequest{Command: command,ChatId: chatid}
	res, err := c.AuthService(ctx, req)
	if err!=nil{
		log.Fatalln("Error when getting response from auth_server: ",err)
	}
	return res.Message,res.Status
}

func separate_param(str string) (string,string,string) {
	command := ""
	param1 := ""
	param2 := ""

	s := strings.Split(str, " ")

	command = s[0]

	switch len(s) {
	case 2:
		param1 = s[1]
	case 3:
		param1 = s[1]
		param2 = s[2]
	}

	return command, param1, param2


	/*if strings.Contains(str, " "){
		space_index := strings.Index(str," ")
		command := str[0:space_index]
		param1 := str[space_index:]
		space_index = strings.Index(param1," ")
		param2 := param1[space_index:]
		return strings.Trim(command," "), strings.Trim(param1," "), strings.Trim(param2," ")

	}
	return strings.Trim(str," "), "",""*/
}

func sendEmail(email string, message string) (response bool){

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

func getBookFromDatabase(course_id int64, subject string) *database.BigResponse {
	fmt.Println("ClientConnectedToDatabase")
	_ = godotenv.Load("globals.env")
	port := os.Getenv("DATABASE_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c:= database.NewDatabaseAccessServiceClient(conn)
	ctx := context.Background()
	req := &database.BookRequest{CourseId: course_id, Subject: subject}
	response, _ := c.GetBooks(ctx,req)
	return response
}

func registerUSer(chat_id string, email string, vkey string) (*database.RegisterResponse,error){
	fmt.Println("ClientConnectedToDatabase")
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
	req := &database.RegisterRequest{ChatId: chat_id,UserEmail: email, Vkey: vkey}
	response, err := c.RegisterUser(ctx, req)
	return response, err
}


func (s *Server) CommandPack(ctx context.Context,req *main_server.CommandPackRequest) (*main_server.CommandPackResponse, error) {
	command:=req.Command
	param1 := ""
	param2 := ""
	command,param1,param2 = separate_param(command)
	
	res := &main_server.CommandPackResponse{Status: false, Response: "Unknown command, use /guide to see available commands"}
	auth_message, auth_status := Auth_service(req.ChatId,command)


	if auth_status && auth_message=="Allowed" {
		fmt.Println(command, "Param1: " + param1, "Param2: " +param2)
		switch command {
		case "/get":
			if param1 == "" {
				res.Status = true
				res.Response = "the Course parameter is empty"
				break
			}
			if param1 != "" && param2 == "" {
				res.Status = true
				course_id, _ := strconv.ParseInt(param1, 10, 32)
				response_books := getBookFromDatabase(course_id, "")
				books_array := response_books.BookPacks

				if len(books_array) == 0 {
					res.Response = "wrong index"
					break
				}

					res.Response = ""
				for _, s := range books_array {
					res.Response = res.Response + s.Subject + " " + s.BookName + ": " + s.BookLink + "\n"
				}
				break
			}
			/*if param1 != "" && param2 != "" {
				res.Status = true
				course_id, _ := strconv.ParseInt(param1, 10, 32)
				subject := param2
				response_books := getBookFromDatabase(course_id, subject)
				books_array := response_books.BookPacks
				res.Response = ""
				for _, s := range books_array {
					res.Response = res.Response + s.Subject + " " + s.BookName + ": " + s.BookLink + "\n"
				}
				break
			}*/

		case "/start":
			res.Status = true
			res.Response = "Hello, use /guide to see all commands"
		case "/guide":
			res.Status = true
			res.Response = "Available commands: /guide, /start, /reg <example_email@astaneit.edu.kz> \n" +
				"/get <your course>"

		case "/reg":
			if param1 == "" {
				res.Status = true
				res.Response = "Email field is empty, use /guide to see available commands"
				break
			}
			if strings.Contains(param1, "@astanait.edu.kz") {
				res.Status = true
				email_address := param1
				vkey := md5.Sum([]byte(email_address))

				registration, _ := registerUSer(req.ChatId, email_address, hex.EncodeToString(vkey[:]))

				if registration.Status == true && registration.Message == "UserAlreadyExist" {
					res.Response = "You have already registered in the system!"
					break
				}
				if registration.Status == true && registration.Message == "UserCreated" {
					email_message := "Hello, dear user! " +
						"\nYou see this message because you have registered your email in Monjjubot. " +
						"\nTo finish registration please follow this link: \n" +
						s.server_domain_name + ":" + s.website_server_port + "/verify/" + hex.EncodeToString(vkey[:]) + "\n" +
						"\n If you did not do that, please close this letter and do not respond. "
					sending_status := sendEmail(email_address, email_message)
					if sending_status {
						res.Response = "Email accepted, please check your mail box! P.s Check spam folder too)"
					} else {
						res.Response = "Error occurred while sending email, pls connect to devs"
					}
				} else {
					res.Response = "Error occurred while registering, pls connect to devs"
				}

			} else {
				res.Status = true
				res.Response = "Email is invalid, or do not belongs to @astanait.edu.kz domen"
			}
		}
	} else {
		res.Status=true
		res.Response="You are not registered"
		return res, nil
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
	main_server.RegisterTelegramBotServiceServer(s, &Server{
		server_address: address,
		website_server_port: os.Getenv("WEBSITE_PORT"),
		server_domain_name: os.Getenv("SERVER_DOMAIN_NAME"),
	})
	log.Println("Command handler Server is running on port: "+port)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}