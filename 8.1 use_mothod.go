package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating,'f', 4, 64)
	return m.Name + ", " + r
}

func main()  {
	m := Movie{
		Name:"Spiderman",
		Rating:3.554345,
	}
	fmt.Println(m.summary())
}