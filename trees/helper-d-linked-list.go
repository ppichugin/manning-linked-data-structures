package main

// *** DoublyLinkedList code ***

type Cell struct {
	data       *Node
	next, prev *Cell
}

type DoublyLinkedList struct {
	topSentinel, bottomSentinel *Cell
}

// Make a new DoublyLinkedList and initialize its sentinels.
func makeDoublyLinkedList() DoublyLinkedList {
	list := DoublyLinkedList{}

	topCell := Cell{nil, nil, nil}
	bottomCell := Cell{nil, nil, nil}

	list.topSentinel = &topCell
	list.topSentinel.next = &bottomCell
	list.topSentinel.prev = nil

	list.bottomSentinel = &bottomCell
	list.bottomSentinel.prev = list.topSentinel
	list.bottomSentinel.next = nil
	return list
}

// Add a cell after me.
func (me *Cell) addAfter(after *Cell) {
	after.next = me.next
	after.prev = me

	me.next.prev = after
	me.next = after
}

// Add a cell before me.
func (me *Cell) addBefore(before *Cell) {
	me.prev.addAfter(before)
}

// Delete the cell and return it.
func (me *Cell) delete() *Cell {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me
}

// Return true if the stack is empty, false otherwise.
func (stack *DoublyLinkedList) isEmpty() bool {
	return stack.topSentinel.next == stack.bottomSentinel
}
