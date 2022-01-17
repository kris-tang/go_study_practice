package main

import "fmt"

func sumNumbers(numbers...int) int{
	total :=0
	for _,number := range numbers{
		total += number
	}
	return total
}

func main(){
	result := sumNumbers(9,8,7,6)
	fmt.Println("The result is %v\n", result)
}