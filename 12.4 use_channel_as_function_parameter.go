package main

import "fmt"

func channelReader(messages <-chan string){
	msg := <-messages
	fmt.Println(msg)
}

func channelWriter(messages chan<- string)  {
	messages <- "Hello world"
}

func channelReaderWriter(messages chan string)  {
	msg := <-messages
	fmt.Println(msg)
	messages <- "Hello world"

}