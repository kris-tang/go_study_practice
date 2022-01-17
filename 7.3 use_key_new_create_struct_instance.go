package main

import "fmt"

type Movie struct{
	Name string
	Rating float32
}

func main()  {
	 m := new(Movie)
	 m.Name = "Metroplois"
	 m.Rating = 0.8837
	 fmt.Println("%+v\n", m)
}