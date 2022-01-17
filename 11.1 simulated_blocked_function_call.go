package main

import (
	"fmt"
	"time"
)

func slowSleep(){
	time.Sleep(time.Second * 7)
	fmt.Println("Sleeper() finished")
}

func main()  {
	slowSleep()
	fmt.Println("I am not shown until slowSleep() completes")
}