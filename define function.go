package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func main(){
	fmt.Println("%v\n", isEven(99))
	fmt.Println("%v\n", isEven(19384))
}