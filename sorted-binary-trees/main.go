package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func (node *Node) insertValue(newData string) {
	newNode := Node{newData, nil, nil}
	for {
		if newData <= node.data {
			if node.left == nil {
				node.left = &newNode
				return
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				node.right = &newNode
				return
			} else {
				node = node.right
			}
		}
	}
}

func (node *Node) findValue(value string) *Node {
	for node != nil {
		if value == node.data {
			return node
		} else if value < node.data {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (node *Node) inorder() string {
	if node == nil {
		return ""
	}
	s := node.left.inorder()
	s += node.data + " "
	s += node.right.inorder()
	return s
}

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insertValue("I")
	root.insertValue("G")
	root.insertValue("C")
	root.insertValue("E")
	root.insertValue("B")
	root.insertValue("K")
	root.insertValue("S")
	root.insertValue("Q")
	root.insertValue("M")

	// Add F.
	root.insertValue("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		_, _ = fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.findValue(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
