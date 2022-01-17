package main

import "fmt"

func main()  {
	name, role := "Richard Jupp", "Darummer"
	err := fmt.Errorf("The %v %v quit", name, role)
	if err != nil {
		fmt.Println(err)
	}
}