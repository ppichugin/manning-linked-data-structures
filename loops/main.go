package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func makeLinkedList() *LinkedList {
	return &LinkedList{&Cell{"SENTINEL", nil}}
}

// addAfter adds a cell after 'm'.
func (m *Cell) addAfter(after *Cell) {
	after.next = m.next
	m.next = after
}

// deleteAfter deletes a cell after 'm'.
func (m *Cell) deleteAfter() *Cell {
	if m.next == nil {
		panic("deleteAfter called on a cell with no next cell")
	}
	deleted := m.next
	m.next = m.next.next
	return deleted
}

func (list *LinkedList) addRange(values []string) {
	// loop over the list to find the last cell
	lastCell := list.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}

	for _, value := range values {
		lastCell.addAfter(&Cell{value, nil})
		lastCell = lastCell.next
	}
}

func (list *LinkedList) toString(separator string) string {
	result := ""
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		if cell.next == nil {
			separator = ""
		}
		result += cell.data + separator
	}
	return result
}

func (list *LinkedList) length() int {
	result := 0
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		result++
	}
	return result
}

func (list *LinkedList) isEmpty() bool {
	return list.sentinel.next == nil
}

func (list *LinkedList) contains(val string) bool {
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		if cell.data == val {
			return true
		}
	}
	return false
}

func (list *LinkedList) find(c Cell) *Cell {
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		if cell.data == c.data {
			return cell
		}
	}
	return nil
}

func (list *LinkedList) remove(c Cell) bool {
	for cell := list.sentinel; cell.next != nil; cell = cell.next {
		if cell.next.data == c.data {
			cell.deleteAfter()
			return true
		}
	}
	return false
}

func (list *LinkedList) removeAt(index int) bool {
	if index < 0 || index >= list.length() || list.isEmpty() {
		return false
	}
	prev := list.sentinel
	for i, cell := 1, list.sentinel.next; cell.next != nil; i, cell = i+1, cell.next {
		if i == index {
			prev.deleteAfter()
			return true
		}
		prev = cell
	}
	return false
}

// append adds a new value to the end of the list
func (list *LinkedList) append(c Cell) {
	// find the last cell
	lastCell := list.sentinel
	for lastCell.next != nil {
		lastCell = lastCell.next
	}
	lastCell.addAfter(&c)
}

// toSlice returns a slice containing the list’s values
func (list *LinkedList) toSlice() []string {
	ll := list.length()

	result := make([]string, 0, ll)
	for cell := list.sentinel.next; cell != nil; cell = cell.next {
		result = append(result, cell.data)
	}
	return result
}

func (list *LinkedList) toStringMax(s string, max int) string {
	result := ""
	for cell := list.sentinel.next; cell != nil && max > 0; cell, max = cell.next, max-1 {
		if cell.next == nil {
			s = ""
		}
		result += cell.data + s
	}
	return result
}

func (list *LinkedList) hasLoop() bool {
	if list.isEmpty() {
		return false
	}
	slow := list.sentinel.next
	fast := list.sentinel.next.next
	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := makeLinkedList()
	list.addRange(values)

	fmt.Println(list.toString(" "))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.toStringMax(" ", 10))
	if list.hasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
