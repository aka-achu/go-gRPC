package client

import (
	"fmt"
	"log"
	"time"
)

func logger(format string, a ...interface{}) {
	log.Printf("CLIENT:\t"+format+"\n", a...)
}

func fatalLogger(format string, a ...interface{}) {
	log.Fatalf("CLIENT: ERROR\t"+format+"\n", a...)
}

func printer(format string, a ...interface{}) {
	fmt.Printf(time.Now().Format("2006/01/02 15:04:05") +" CLIENT:\t"+format+"\n", a...)
}
