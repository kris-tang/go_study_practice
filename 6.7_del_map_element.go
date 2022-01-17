package main

import "fmt"

func main()  {
	var players = make(map[string]int)
	players["cook"] = 38575
	players["briostow"] = 35565
	players["stokes"] = 385485
	delete(players, "cook" )
	fmt.Println(players)
}