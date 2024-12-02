package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// InfoLogger is a logger instance which prints to stdout
	InfoLogger *log.Logger
	// ErrorLogger is a logger instance which prints to stderr
	ErrorLogger *log.Logger
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A simple go application says hello!")

}

func main() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Lmsgprefix)
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	InfoLogger.Printf(`Listening and print on port: %s`, port)
	http.HandleFunc("/", handler)
	InfoLogger.Printf(`Info log: Handler started and running!`)
	http.ListenAndServe(":"+port, nil)
}
