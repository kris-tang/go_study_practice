package main

import (
	"fmt"
	"time"
)

func pinger(c chan string)  {
	t := time.NewTicker(3 * time.Second)
	for{
		c <-"ping"
		<-t.C
	}
}

func main()  {
	messages := make(chan string)
	go pinger(messages)
	for i :=5; i>0; i-- {
		msg := <-messages
		fmt.Println(msg)
	}
}