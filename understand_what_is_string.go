package main

import "fmt"

func main()  {
	var s1 string = "hello"
	fmt.Println(len(s1))
	var s2 string = "hello"
	fmt.Println(s2[0])
	var s3 string = "hello"
	fmt.Println("%q", s3[0])
	fmt.Println("%b", s3[0])
	var s4 string = "にほんご"
	fmt.Println(len(s4))
}
