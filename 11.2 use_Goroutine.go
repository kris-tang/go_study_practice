package main

import (
	"fmt"
	"time"
)

func slowFunc()  {
	time.Sleep(time.Second * 2)
	fmt.Println("Sleeper() finished")
}

func main()  {
	go slowFunc()
	fmt.Println("I am now shown straightaway")
}