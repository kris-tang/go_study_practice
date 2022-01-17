package main

import "time"

messages := make(chan string)
stop := make(chan bool)

for{
	select{
	case <-stop:
		return
	case msg := <-messages:
		fmt.Println(msg)
}
}

func sender(c chan string){
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

messages := make(chan string)
go sender(messages)

for{
	select {
	case msg := <-messages:
		fmt.Println(msg)
	}
}

stop := make(chan bool)

for{
	select{
	case <- stop:
		return
	case msg := <-messages:
		fmt.Println(msg)
}
}