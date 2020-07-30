package server

import (
	"fmt"
	"log"
	"time"
)

func logger(format string, a ...interface{}) {
	log.Printf("SERVER:\t"+format+"\n", a...)
}

func fatalLogger(format string, a ...interface{}) {
	log.Fatalf("SERVER: ERROR\t"+format+"\n", a...)
}

func printer(format string, a ...interface{}) {
	fmt.Printf(time.Now().Format("2006/01/02 15:04:05") +" SERVER:\t"+format+"\n", a...)
}

