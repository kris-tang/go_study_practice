package main

import "fmt"

func sayHi()(x,y string){
	x = "Hello"
	y = "world"
	return
}

func main(){
	fmt.Println(sayHi())
}