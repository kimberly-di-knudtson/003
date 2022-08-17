package main

import (
	"fmt"
)

type Node struct {
	data string
	left *Node
	right *Node
}

//the tree
func main() {
	root := new Node("root", Node("left", Node("left")), Node("right"))

}