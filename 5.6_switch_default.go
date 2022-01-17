package main

import "fmt"

func main(){
	s := "djjg"
	switch s{
	case "c":
		fmt.Println("this is c")
	case "e":
		fmt.Println("this is e")
	default:
		fmt.Println("I do not recongnise the letter")
	}
}