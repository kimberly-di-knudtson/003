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

//Make some rules...
//start at leftmost leaf or rightmost leaf - doesn't matter - choose leftmost
//at a node serialize(node) gives ("self","left","right")
//serialize(root.left.left) = "(left.left,nil,nil)" = this.val+serialize(this.left)+serialize(this.right)
//serialize(root.left) = "(left,(left.left,nil,nil),nil" = this.val+serialize(this.left)+serilize(this.right)
//serialize(root.right) = "(right,nil,nil)" =this.val+serialize(this.left)+serialize(this.right)
//serialize(root) = "(root,(left,(left.left,nil,nil),nil),(right,nil,nil))"

func serialize(node Node) string {
	s := node.val + ","
	if node.left != nil {
		s = s + serialize(*node.left) + ","
	} else {
		s = s + "nil" + ","
	}
	if node.right != nil {
		s = s + serialize(*node.right)
	} else {
		s = s + "nil"
	}
	return "(" + s + ")"
}

//root := Node{val: "root"}
//root.left = &Node{val: "left"}
//root.left.left = &Node{val: "left.left"}
//root.right = &Node{val: "right"}
//serialize(root) = "(root,(left,(left.left,nil,nil),nil),(right,nil,nil))"

//ds("(root,(left,(LEFTNODE),nil),(right,nil,nil))", leftnode)
//ds("(root,(LEFTNODE),(right,nil,nil))")
func deserialize(s string, leftnode Node, rightnode Node) Node {
	//format is (self,left,right)
}

func main() {
	s := "root"
	root := deserialize(s)
	fmt.Print(s)

}*/
