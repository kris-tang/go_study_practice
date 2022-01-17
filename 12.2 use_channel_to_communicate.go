package main

import "time"

func slowFunc(c chan string)  {
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finshed"
}

func main()  {
	c := make(chan string)
	go slowFunc(c)

	msg := <-c
	println(msg)
}