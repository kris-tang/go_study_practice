package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("plese input a value")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	var s2 int64 = int64(s1)
	fmt.Println("read data", s1)
	if s2 > 5{
		fmt.Println("s gt 5")
	}

}
