package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (b *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (c *R2D2) PowerOn() error {
	if c.Broken{
		return errors.New("R2D2 Robot is broken!")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main()  {
	t := T850{
		Name:"The Terminator",
	}
	r := R2D2{
		Broken: true,
	}

	err :=Boot(&r)
	if err !=nil{
		fmt.Println(err)
	} else {
		fmt.Println("R2D2 Robot is Powered on")
	}

	err =Boot(&t)
	if err !=nil{
		println(err)
	} else {
		fmt.Println("T850 Robot is Powered on")
	}

}