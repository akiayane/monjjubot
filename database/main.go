package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func main() {
	dsn := flag.String("dns", "postgres://postgres:root@localhost:5432/telebot", "Postgre data source name")

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal("ERROR for openDB func: ", err)
	}

	defer db.Close()

	var greeting string
	err = db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	model := SnippetModel{DB: db}
	books, _ := model.getByCourse(2)

	for _, book := range books {
		fmt.Println(book)
	}
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		return nil, err
	}
	return conn, nil
}
