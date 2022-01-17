package main

import (
	"fmt"
	"time"
)

func sender(c chan string)  {
	t := time.NewTicker(1 * time.Second)
	for{
		c <- "I'm sending a message"
		<-t.C
	}
}

func main()  {
	messages := make(chan string)
	quit := make(chan bool)

	go sender(messages)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("time's up")
		quit <- true
	}()

	for{
		select {
		case <-quit:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}