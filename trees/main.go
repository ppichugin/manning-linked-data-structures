package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data        string
	left, right *Node
}

func (node *Node) displayIndented(indent string, depth int) string {
	if node == nil {
		return ""
	}
	s := strings.Repeat(indent, depth) + node.data + "\n"
	s += node.left.displayIndented(indent, depth+1)
	s += node.right.displayIndented(indent, depth+1)
	return s
}

func (node *Node) preorder() string {
	if node == nil {
		return ""
	}
	s := node.data + " "
	s += node.left.preorder()
	s += node.right.preorder()
	return s
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

func (node *Node) postorder() string {
	if node == nil {
		return ""
	}
	s := node.left.postorder()
	s += node.right.postorder()
	s += node.data + " "
	return s
}

func (node *Node) breadthFirst() string {
	queue := makeDoublyLinkedList()
	queue.enqueue(node)
	s := ""
	for !queue.isEmpty() {
		node := queue.dequeue()
		s += node.data + " "
		if node.left != nil {
			queue.enqueue(node.left)
		}
		if node.right != nil {
			queue.enqueue(node.right)
		}
	}
	return s
}

func buildTree() *Node {
	aNode := Node{"A", nil, nil}
	bNode := Node{"B", nil, nil}
	cNode := Node{"C", nil, nil}
	dNode := Node{"D", nil, nil}
	eNode := Node{"E", nil, nil}
	fNode := Node{"F", nil, nil}
	gNode := Node{"G", nil, nil}
	hNode := Node{"H", nil, nil}
	iNode := Node{"I", nil, nil}
	jNode := Node{"J", nil, nil}

	aNode.left = &bNode
	aNode.right = &cNode
	bNode.left = &dNode
	bNode.right = &eNode
	eNode.left = &gNode
	cNode.right = &fNode
	fNode.left = &hNode
	hNode.left = &iNode
	hNode.right = &jNode

	return &aNode
}

func main() {
	// Build a tree.
	aNode := buildTree()

	// Display with indentation.
	fmt.Println(aNode.displayIndented("  ", 0))

	// Display traversals.
	fmt.Println("Preorder:\t", aNode.preorder())
	fmt.Println("Inorder:\t", aNode.inorder())
	fmt.Println("Postorder:\t", aNode.postorder())
	fmt.Println("Breadth first:\t", aNode.breadthFirst())
}
