package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/databaseproto"
	"net"
	"os"
)

type Server struct {
	databaseproto.UnimplementedDatabaseAccessServiceServer
	SnippetModel
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		return nil, err
	}
	return conn, nil
}

func(s *Server) GetBooks(ctx context.Context, request *databaseproto.BookRequest) (*databaseproto.BigResponse, error){

	course_id := request.CourseId
	subject := request.Subject

	var books []*databaseproto.BookPack

	err := errors.New("empty error")
	if subject==""{
		books, err = s.getByCourse(int(course_id))
		if err!=nil{
			log.Println("Error occurred while calling getBySubject function on Database_Access_Server", err)
		}
	} else {
		books, err = s.getBySubject(int(course_id), subject)
		if err!=nil{
			log.Println("Error occurred while calling getByCourse function on Database_Access_Server", err)
		}
	}

	bigResponse:= &databaseproto.BigResponse{BookPacks: books}
	return bigResponse, err

}

func(s *Server) RegisterUser(ctx context.Context, request *databaseproto.RegisterRequest) (*databaseproto.RegisterResponse, error){

	chat_id := request.ChatId
	email := request.UserEmail
	vkey := request.Vkey

	res := &databaseproto.RegisterResponse{Status: false, Message: ""}

	if s.userExist(chat_id, email){
		res.Status = true
		res.Message = "UserAlreadyExist"
	}else{
		stat, err := s.registerUser(chat_id, email, vkey)
		if err==nil && stat==true{
			res.Status = true
			res.Message = "UserCreated"
		}
		if err!=nil{
			log.Println("Error occured while registering user on database_access_server", err)
		}
	}

	return res, nil
}

func(s *Server) ConfirmRegister(ctx context.Context, request *databaseproto.ConfirmRequest) (*databaseproto.RegisterResponse, error){

	vkey := request.Vkey

	res := &databaseproto.RegisterResponse{Status: false, Message: "UserNotConfirmed"}

	if s.confirmRegister(vkey){
		res.Status = true
		res.Message = "RegistrationConfirmed"
	}

	return res, nil
}


func main() {
	dsn := flag.String("dns", "postgres://postgres:root@localhost:5432/telebot", "Postgre data source name")

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal("ERROR for openDB func: ", err)
	}

	defer db.Close()

	_ = godotenv.Load("globals.env")
	port := os.Getenv("DATABASE_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")

	/*model := SnippetModel{DB: db}
	books_array, _ := model.getByCourse(1)

	response := ""
	for _,s := range books_array{
		response = response + s.Subject + s.BookLink + s.BookLink + "\n"
	}
	fmt.Println(response)*/

	l, err := net.Listen("tcp", address+":"+port)
	if err!=nil{
		log.Fatalln("Error occurred while deploying the server",err)
	}
	s := grpc.NewServer()
	model := SnippetModel{DB: db}
	databaseproto.RegisterDatabaseAccessServiceServer(s, &Server{SnippetModel: model})
	log.Println("Database_Access_Server is running on port: "+port)
	if err := s.Serve(l); err!=nil{
		log.Fatalln("Error occured while listening for Database_Access_Server", err)
	}


}

