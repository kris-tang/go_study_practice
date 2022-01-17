package main

import "fmt"

var s = "Hello World"
func main(){
	fmt.Println("Print 's' varibale from outer block %v\n", s)
	b := true
	if b{
		i :=445
		fmt.Println("Print 'b' variable from outer block %v\n", b)
		if b != false{
			fmt.Println("Print 'i' variable from outer block %v\n", i)
		}
	}
}