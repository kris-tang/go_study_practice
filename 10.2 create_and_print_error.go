package main

import (
	"errors"
	"fmt"
)

func main()  {
	var err error
	err = errors.New("something went wrong !!!!!")
	if err != nil {
		fmt.Println(err)
	}
}