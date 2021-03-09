package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"monjjubot/databaseproto"
	"net/http"
	"os"
)

func verifyOnDatabase(vkey string) (bool,string){
	fmt.Println("ClientConnected")
	_ = godotenv.Load("globals.env")
	port := os.Getenv("DATABASE_SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := databaseproto.NewDatabaseAccessServiceClient(conn)
	ctx := context.Background()
	req := &databaseproto.ConfirmRequest{Vkey: vkey}

	res, err := c.ConfirmRegister(ctx, req)
	if err!=nil{
		log.Fatalln("Error when getting response from database_access_server: ",err)
	}
	return res.Status, res.Message
}

func homepage(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
}

func verify(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	vkey := vars["id"]
	confirmation_status, confirmation_message := verifyOnDatabase(vkey)
	if confirmation_status==true && confirmation_message=="RegistrationConfirmed"{
		_, _ = fmt.Fprintf(w, "The account is now verified with vkey: "+vkey)
	}else{
		_, _ = fmt.Fprintf(w, "Sorry, error occurred while verification, connect to devs. Vkey is : "+vkey)
		log.Fatalln("Error occurred while verification")
	}
}

func handleRequests(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homepage)
	router.HandleFunc("/verify/{id}", verify).Methods("GET")
	_ = godotenv.Load("globals.env")
	port := os.Getenv("WEBSITE_PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main(){
	log.Println("Website is runnning ")
	handleRequests()
}