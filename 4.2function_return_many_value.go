package main

import "fmt"

func getPrize()(int, string){
	i :=2
	s := "goldfish"
	return i, s
}
func main(){
	quantity, prize :=getPrize()
	fmt.Println("You won %d %s\n", quantity,prize)
}
