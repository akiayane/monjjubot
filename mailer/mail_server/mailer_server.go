package main

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/mailer"
	"net"
	"net/smtp"
	"os"
)

type Server struct {
	mailer.UnimplementedEmailSendingServiceServer
}


func(s *Server) SendEmail(ctx context.Context,request *mailer.EmailRequest) (*mailer.EmailResponse, error) {
	email_address := request.Email
	email_message := request.Message
	from := os.Getenv("GMAIL")
	password := os.Getenv("GMAIL_PASSWORD")
	to := []string{
		email_address,
	}
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_POST")

	message := []byte(email_message)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	response := &mailer.EmailResponse{Status: true}
	if err != nil {
		response.Status = false
		log.Fatalln("Error occured while sending email", err)
	}
	return response,nil
}

func main(){
	_ = godotenv.Load("globals.env")
	port := os.Getenv("MAILER_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	l, err := net.Listen("tcp", address+":"+port)
	if err!=nil{
		log.Fatalln("Error occurred while deploying the server",err)
	}
	s := grpc.NewServer()
	mailer.RegisterEmailSendingServiceServer(s, &Server{})
	log.Println("Mailer_Server is running on port: "+port)
	if err := s.Serve(l); err!=nil{
		log.Fatalln("Error occured while listening for mailer server", err)
	}
}
