package main

import (
	routes "css/mlc/routes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := "127.0.0.1:1234"
	runServer(port)
}

func runServer(port string) {
	r := routes.Start_Routes()
	fmt.Printf("Listening on port %s\n", port)
	server := http.Server{
		Addr:    port,
		Handler: r,
	}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
