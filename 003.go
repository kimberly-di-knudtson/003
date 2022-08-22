package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

//the tree
func main() {
	root := Node{val: "root"}
	//root.left = &Node{val: "left"}
	//root.left.left = &Node{val: "left.left"}
	root.right = &Node{val: "right"}
	s := serialize(root)
	fmt.Println(s)
	sz_root, err := deserialize(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(serialize(*sz_root))
	}

	//s = "(root,left,right)"
	//sz, _ := deserialize(s)
	//fmt.Print(sz)

}

//node = Node('root', Node('left', Node('left.left')), Node('right'))
//assert deserialize(serialize(node)).left.left.val == 'left.left'

//Make some rules...
//start at leftmost leaf or rightmost leaf - doesn't matter - choose leftmost
//at a node serialize(node) gives ("self","left","right")
//serialize(root.left.left) = "(left.left,,)" = this.val+serialize(this.left)+serialize(this.right)
//serialize(root.left) = "(left,(left.left,,)," = this.val+serialize(this.left)+serilize(this.right)
//serialize(root.right) = "(right,,)" =this.val+serialize(this.left)+serialize(this.right)
//serialize(root) = "(root,(left,(left.left,,),),(right,,))"

func serialize(node Node) string {
	s := node.val + ","
	if node.left != nil {
		s = s + serialize(*node.left) + ","
	} else {
		s = s + ","
	}
	if node.right != nil {
		s = s + serialize(*node.right)
	}
	return "(" + s + ")"
}

//deserialize a string enclosed in (), format (self,left,right)
func deserialize(s string) (*Node, error) {
	n := new(Node)
	//check that we have an enclosure
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		end_val := strings.Index(s, ",")
		n.val = s[1:end_val]
		fmt.Printf("n.val = %v\n", n.val)
		next_token := s[end_val+1 : end_val+2] //s[end_val] == "," - next is either empty, unenclosed or enclosed
		if next_token == "," {                 //no left node
			end_val = end_val + 1
		} else if next_token == "(" { //find the end of this parenthesized enclosure and deserialize it, it is left node
			left, err := get_enclosure(s[end_val+1:])
			if err != nil {
				return n, err
			}
			left_node, err := deserialize(left)
			if err != nil {
				return n, err
			}
			n.left = left_node
			end_val = end_val + len(left) + 1
		} else {
			end_left_val := strings.Index(s[end_val+1:], ",")
			left := Node{val: s[end_val+1 : end_left_val]}
			n.left = &left
			end_val = end_left_val + 1
		}
		//next_token = s[end_val : end_val+1]
		next_token = s[end_val+1 : end_val+2]
		if next_token == ")" { //no right node, return n
			return n, nil
		} else if next_token == "(" { //find the end of this parenthesized enclosure and deserialize it, it is left node
			right, err := get_enclosure(s[end_val+1:])
			if err != nil {
				return n, err
			}
			right_node, err := deserialize(right)
			if err != nil {
				return n, err
			}
			n.right = right_node
			return n, nil
		} else {
			end_right_val := strings.Index(s[end_val+1:], ")")
			right := Node{val: s[end_val+1 : end_right_val]}
			n.right = &right
			return n, nil
		}
	} else {
		e := fmt.Errorf("deserialize string \",%v,\" is not enclosed in ()", s)
		return n, e
	}
}

func get_enclosure(s string) (string, error) {
	enclosed_string := ""
	offset := 0
	if strings.HasPrefix(s, "(") {
		for counter := 1; counter > 0; {
			offset = offset + 1
			next_open := strings.Index(s[offset:], "(")
			next_close := strings.Index(s[offset:], ")")
			if next_open > next_close || next_open == -1 {
				counter = counter - 1
				offset = offset + next_close
			} else if next_open < next_close && next_open != -1 {
				counter = counter + 1
				offset = offset + next_open
			}
		}
		if offset == len(s)-1 {
			enclosed_string = s
		} else {
			enclosed_string = s[:offset+1] //slice doesn't include last index's value
		}
	} else {
		return "", fmt.Errorf("\"%v\" is not an enclosed string", s)
	}
	return enclosed_string, nil
}
