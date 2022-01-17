package main

import "fmt"

func main()  {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogne"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	var smellycheeses = make([]string, 2)
	copy(smellycheeses, cheeses)
	fmt.Println(len(smellycheeses))
	fmt.Println(smellycheeses)
}