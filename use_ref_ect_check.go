package main

import (
	"fmt"
	"reflect"
)

func main(){
	var s string = "string"
	var i int = 2
	var f float32 = 1.2222

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}