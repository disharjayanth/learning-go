package main

import "fmt"

type Node struct {
	val         int
	left, right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		return &Node{
			val:   val,
			left:  nil,
			right: nil,
		}
	}
	if val < n.val {
		n.left = n.left.Insert(val)
	}
	if val > n.val {
		n.right = n.right.Insert(val)
	}
	return n
}

func (n *Node) Contains(val int) bool {
	switch {
	case n == nil:
		return false
	case val < n.val:
		return n.left.Contains(val)
	case val > n.val:
		return n.right.Contains(val)
	default:
		return true
	}
}

func main() {
	var n *Node
	n = n.Insert(5)
	n = n.Insert(3)
	n = n.Insert(10)
	n = n.Insert(2)
	fmt.Println(n.Contains(2))
	fmt.Println(n.Contains(12))
}
