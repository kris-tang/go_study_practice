package main

import "fmt"

func main(){
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers{
		fmt.Println("the index of the loop is:", i)
		fmt.Println("the value of the array is:", n)
	}
}