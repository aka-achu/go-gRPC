package main

import (
	"github.com/aka-achu/go-gRPC/client"
	"github.com/aka-achu/go-gRPC/server"
	"github.com/subosito/gotenv"
	"log"
)

func init() {
	if gotenv.Load(".env") != nil {
		log.Fatal("Failed to load the env file")
	}
}
func main() {
	go server.Initialize()
	client.Initialize()
}
