package main

import "fmt"

func anotherFuncation(f func() string) string{
	return f()
}

func main(){
	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFuncation(fn))
}