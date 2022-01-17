package main

import "fmt"

type Movie struct{
	Name string
	Rating  float32
}

func main()  {
	m := Movie{
		Name: "Metroplois",
		Rating: 0.3594,
	}
	fmt.Println("%+v\n", m)
}