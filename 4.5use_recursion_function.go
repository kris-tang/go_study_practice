package main

import "fmt"

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 8{
		fmt.Println("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
		fmt.Println("I'm hungry! I've eaten %d\n", eaten)
		return feedMe(portion,eaten)
}

func main(){
	fmt.Println(feedMe(1,0))
}