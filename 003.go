package main

import (
	"fmt"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

//the tree
func main() {
	root := Node{val: "root"}
	root.left = &Node{val: "left"}
	root.left.left = &Node{val: "left.left"}
	root.right = &Node{val: "right"}
	s := serialize(root)
	fmt.Print(s)

}

//node = Node('root', Node('left', Node('left.left')), Node('right'))
//assert deserialize(serialize(node)).left.left.val == 'left.left'

func serialize(node Node) string {
	return node.left.left.val
}
