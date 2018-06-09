package main

import "fmt"

// Program is WIP.

type Planet struct {
	Size   int
	Radius int
}

func main() {
	world := Planet{1, 2}
	fmt.Println(world.Size)
}
