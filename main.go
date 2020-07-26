package main

import (
	"github.com/aka-achu/go-gRPC/client"
	"github.com/aka-achu/go-gRPC/server"
)

func main() {
	go server.Initialize()
	client.Initialize()
}
