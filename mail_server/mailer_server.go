package main

import (
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

func sendEmail(request mailer.EmailRequest) mailer.EmailResponse {
	param := request.Email
	from := os.Getenv("GMAIL")
	password := os.Getenv("GMAIL_PASSWORD")
	to := []string{
		param,
	}
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_POST")

	message := []byte("Test, hello "+param)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	response := mailer.EmailResponse{Status: true}
	if err != nil {
		response.Status = false
		log.Fatalln("Error occured while sending email", err)
	}
	return response
}

func main(){
	_ = godotenv.Load("globals.env")
	port := os.Getenv("MAILER_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	l, err := net.Listen("tcp", address+port)
	if err!=nil{
		log.Fatalln("Error occurred while deploying the server",err)
	}
	s := grpc.NewServer()
	mailer.RegisterEmailSendingServiceServer(s, &Server{})
	if err := s.Serve(l); err!=nil{
		log.Fatalln("Error occured while listening for mailer server", err)
	}
}
