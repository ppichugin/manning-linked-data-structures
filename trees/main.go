package main

import (
	"fmt"
)

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
