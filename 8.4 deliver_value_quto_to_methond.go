package main

import "fmt"

type Triangle struct{
	base float64
	height float64
}

func (t *Triangle) changeBase(f float64)  {
	t.base = f
}

func main()  {
	t := Triangle{base:3, height: 6}
	t.changeBase(6)
	fmt.Println(t.base)
}