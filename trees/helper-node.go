package main

import "strings"

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
