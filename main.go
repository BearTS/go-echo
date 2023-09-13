package main

import (
	"github.com/BearTS/go-echo/cmd"
	"github.com/BearTS/go-echo/pkg/logger"
	"github.com/joho/godotenv"
)

// Message represents the message structure you expect to send to the RabbitMQ queue.

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal(err)
		return
	}
	cmd.Execute()
}
