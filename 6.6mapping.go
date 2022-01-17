package main

import "fmt"

func main()  {
	var players = make(map[string]int)
	players["cook"] = 35
	players["baristow"] = 34566
	players["stokes"] = 34856
	fmt.Println(players["cook"])
	fmt.Println(players["stokes"])
}