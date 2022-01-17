package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var i int = 1
	var s string = " egg"
	intTostring := strconv.Itoa(i)
	var breakfast string = intTostring + s
	fmt.Println(breakfast)
}