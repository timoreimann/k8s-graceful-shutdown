package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const waitTime = 5 * time.Second

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Starting")
	<-c
	fmt.Printf("Waiting %s before terminating\n", waitTime)
	time.Sleep(waitTime)
	fmt.Println("Terminating")
}
