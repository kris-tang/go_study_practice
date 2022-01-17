package main

import "fmt"

type Movie struct{
	Name string
	Rating float32
}

func main()  {
	var m Movie
	fmt.Printf("%+v\n", m)
	m.Name = "Metroplois"
	m.Rating = 0.998178
	fmt.Printf("%+v\n", m)
}