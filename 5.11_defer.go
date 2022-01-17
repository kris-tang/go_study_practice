package main

import "fmt"

func main(){
	defer fmt.Println("I am the firest defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello World")
}