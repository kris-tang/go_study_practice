package main

import "fmt"

type Movie struct{
	Name  string
	Rating float32
}

func main()  {
	m := Movie{
		Name: "Citizen Kan",
		Rating: 10,
	}
	fmt.Printf(m.Name, m.Rating)
}