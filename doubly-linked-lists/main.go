package main

import "fmt"

type Cell struct {
	prev *Cell
	data string
	next *Cell
}

type DoublyLinkedList struct {
	topSentinel    *Cell
	bottomSentinel *Cell
}

func makeDoublyLinkedList() *DoublyLinkedList {
	list := new(DoublyLinkedList)
	list.topSentinel = new(Cell)
	list.bottomSentinel = new(Cell)
	list.topSentinel.next = list.bottomSentinel
	list.bottomSentinel.prev = list.topSentinel
	return list
}

// addAfter adds a cell after 'm'.
func (m *Cell) addAfter(after *Cell) {
	after.next = m.next
	after.prev = m
	m.next.prev = after
	m.next = after
}

// addBefore adds a cell before 'm'.
func (m *Cell) addBefore(before *Cell) {
	pr := m.prev
	pr.addAfter(before)
}

// delete removes the current cell.
func (m *Cell) delete() *Cell {
	m.prev.next = m.next
	m.next.prev = m.prev
	return m
}

func (list *DoublyLinkedList) addRange(values []string) {
	for _, value := range values {
		cell := Cell{
			data: value,
		}
		list.bottomSentinel.addBefore(&cell)
	}
}

func (list *DoublyLinkedList) toString(separator string) string {
	result := ""
	for cell := list.topSentinel.next; cell != list.bottomSentinel; cell = cell.next {
		if cell.next == list.bottomSentinel {
			separator = ""
		}
		result += cell.data + separator
	}
	return result
}

func (list *DoublyLinkedList) length() int {
	result := 0
	for cell := list.topSentinel.next; cell != list.bottomSentinel; cell = cell.next {
		result++
	}
	return result
}

// queue functions

// Add an item to the top of the queue.
func (queue *DoublyLinkedList) enqueue(value string) {
	cell := Cell{
		data: value,
	}
	queue.topSentinel.addAfter(&cell)
}

// Remove an item from the bottom of the queue.
func (queue *DoublyLinkedList) dequeue() string {
	cell := queue.bottomSentinel.prev.delete()
	return cell.data
}

// Add an item at the bottom of the deque.
func (deque *DoublyLinkedList) pushBottom(value string) {
	cell := Cell{
		data: value,
	}
	deque.bottomSentinel.addBefore(&cell)
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) pushTop(value string) {
	cell := Cell{
		data: value,
	}
	deque.topSentinel.addAfter(&cell)
}

// Remove an item from the top of the deque.
func (deque *DoublyLinkedList) popTop() string {
	cell := deque.topSentinel.next.delete()
	return cell.data
}

// Add an item at the top of the deque.
func (deque *DoublyLinkedList) popBottom() string {
	cell := deque.bottomSentinel.prev.delete()
	return cell.data
}

// isEmpty returns true if the list is empty.
func (list *DoublyLinkedList) isEmpty() bool {
	return list.topSentinel.next == list.bottomSentinel
}

func main() {
	// Make a list from a slice of values.
	//list := makeDoublyLinkedList()
	//animals := []string{
	//	"Ant",
	//	"Bat",
	//	"Cat",
	//	"Dog",
	//	"Elk",
	//	"Fox",
	//}
	//list.addRange(animals)
	//fmt.Println(list.toString(" "))

	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := makeDoublyLinkedList()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.isEmpty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := makeDoublyLinkedList()
	deque.pushTop("Ann")
	deque.pushTop("Ben")
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Cat")
	fmt.Printf("%s ", deque.popBottom())
	fmt.Printf("%s ", deque.popBottom())
	deque.pushBottom("F-Dan")
	deque.pushTop("Eva")
	for !deque.isEmpty() {
		fmt.Printf("%s ", deque.popBottom())
	}
	fmt.Printf("\n")
}
