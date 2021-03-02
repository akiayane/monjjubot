package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func homepage(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
}

func verify(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	vkey := vars["id"]
	_, _ = fmt.Fprintf(w, "This vkey is now verified: "+vkey)
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