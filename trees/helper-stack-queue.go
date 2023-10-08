package main

// *** Stack functions ***

// Push an item onto the top of the list right after the sentinel.
func (stack *DoublyLinkedList) push(node *Node) {
	cell := Cell{data: node}
	stack.topSentinel.addAfter(&cell)
}

// Pop an item off of the list (from right after the sentinel).
func (stack *DoublyLinkedList) pop() *Node {
	if stack.isEmpty() {
		panic("Cannot pop from an empty list")
	}

	return stack.topSentinel.next.delete().data
}

// *** Queue functions ***

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(node *Node) {
	queue.push(node)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() *Node {
	if queue.isEmpty() {
		panic("Cannot dequeue from an empty list")
	}
	return queue.bottomSentinel.prev.delete().data
}
