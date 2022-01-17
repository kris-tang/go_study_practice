package main

import "fmt"

func main()  {
	channel1 := make(chan string)
	channel2 := make(chan string)
	select{
	case msg1 := <-channel1:
	fmt.Println("received", msg1)
	case msg2 := <-channel2:
	fmt.Println("received", msg2)

	}
}
