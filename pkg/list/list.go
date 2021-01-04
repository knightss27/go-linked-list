package list

import (
	"fmt"
)

// Node type - represents internal list item
type node struct {
	data interface{}
	next *node
	prev *node
}

// List type - represents a linked list
type List struct {
	size int
	head *node
	tail *node
}

// NewList - returns a pointer to a new Linked List
func NewList() *List {
	m := new(List)
	return m
}

// Size - returns an integer representing the number of items in the list
func (list *List) Size() int {
	return list.size
}

// GetFirst - returns the data of the first item in the list
func (list *List) GetFirst() interface{} {
	return list.head.data
}

// GetLast - returns the data of the last item in the list
func (list *List) GetLast() interface{} {
	return list.tail.data
}

// AddFirst - adds a new item to the start of the list
func (list *List) AddFirst(data interface{}) *List {

	if list.size == 0 {
		list.head = &node{
			data, nil, nil,
		}
		list.tail = list.head
	} else {
		list.head.prev = &node{
			data: data,
			next: list.head,
			prev: nil,
		}
		list.head = list.head.prev
	}

	list.size++
	return list
}

// AddLast adds a new item to the end of the list
func (list *List) AddLast(data interface{}) *List {

	if list.size == 0 {
		list.head = &node{
			data, nil, nil,
		}
		list.tail = list.head
	} else {
		list.tail.next = &node{
			data, nil, list.tail,
		}
		list.tail = list.tail.next
	}

	list.size++
	return list
}

// RemoveFirst - removes the first item in the list
func (list *List) RemoveFirst() *List {
	if list.size == 0 {
		return list
	}

	if list.size == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.head.next.prev = nil
		list.head = list.head.next
	}

	list.size--
	return list
}

// RemoveLast - removes the last item in the list
func (list *List) RemoveLast() *List {
	if list.size == 0 {
		return list
	}

	if list.size == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
	}

	list.size--
	return list
}

// Clear - clears the list's head and tail
func (list *List) Clear() *List {
	list.head = nil
	list.tail = nil
	list.size = 0
	return list
}

// IsEmpty - returns whether the list has any items in it
func (list *List) IsEmpty() bool {
	return list.size == 0
}

// Add - adds a new item at the specified index
func (list *List) Add(index int, data interface{}) *List {

	if index == 0 {
		return list.AddFirst(data)
	} else if index == list.size {
		return list.AddLast(data)
	}

	if list.checkIndex(index) {
		return list
	}

	toReplace := list.getNode(index)
	toReplace.prev.next = &node{
		data, toReplace, toReplace.prev,
	}
	toReplace.prev = toReplace.prev.next

	list.size++
	return list
}

// Get - gets the item at a spcified index
func (list *List) Get(index int) interface{} {
	if list.checkIndex(index) {
		return ""
	}

	if index == 0 {
		return list.GetFirst()
	} else if index == list.size-1 {
		return list.GetLast()
	}

	return list.getNode(index).data
}

// Set - sets the item at a specified index to the new item
func (list *List) Set(index int, data interface{}) *List {

	if list.checkIndex(index) {
		return list
	}

	toSet := list.getNode(index)
	toSet.data = data

	return list
}

// Remove - removes the item at a specified index
func (list *List) Remove(index int) *List {
	if index == 0 {
		return list.RemoveFirst()
	} else if index == list.size-1 {
		return list.RemoveLast()
	}

	if list.checkIndex(index) {
		return list
	}

	toRemove := list.getNode(index)

	toRemove.prev.next = toRemove.next
	toRemove.next.prev = toRemove.prev

	list.size--
	return list
}

// IndexOf - returns an integer representing the index of the specified item
func (list *List) IndexOf(data interface{}) int {
	counter := 0
	currentNode := list.head
	for currentNode != nil {
		if currentNode.data == data {
			return counter
		}
		currentNode = currentNode.next
		counter++
	}
	return -1
}

// Contains - returns a boolean representing whether the specified item exists in the list
func (list *List) Contains(data interface{}) bool {
	return list.IndexOf(data) != -1
}

// getNode - internal function to fetch a node at a specific index
func (list *List) getNode(index int) *node {
	var toReturn *node

	if index == 0 {
		return list.head
	} else if index == list.size-1 {
		return list.tail
	}

	if index > list.size/2 {
		toReturn = list.tail
		for i := list.size; i > index; i-- {
			toReturn = toReturn.prev
		}
	} else {
		toReturn = list.head
		for i := 0; i < index; i++ {
			toReturn = toReturn.next
		}
	}
	return toReturn
}

// checkIndex - internal function to check index bounds
func (list *List) checkIndex(index int) bool {
	if index < 0 || index >= list.size {
		fmt.Printf("error: index %v out of bounds for list of size %d \n", index, list.size)
		return true
	}
	return false
}

// Iterator type - represents an iterator of a linked list
type Iterator struct {
	hasNext bool
	current *node
	next    *node
}

// Iterator - returns a new iterator associated with the received linked list
func (list *List) Iterator() *Iterator {
	toReturn := Iterator{}

	if list.size >= 1 {
		toReturn.hasNext = true
		toReturn.next = list.head
	}

	return &toReturn
}

// Next - returns a boolean representing whether there is another item to iterate on
func (iterator *Iterator) Next() bool {
	toReturn := iterator.hasNext
	iterator.current = iterator.next

	if iterator.next.next == nil {
		iterator.hasNext = false
	}

	if iterator.hasNext {
		iterator.next = iterator.next.next
	}

	return toReturn
}

// Item - returns the current item being iterated on
func (iterator *Iterator) Item() interface{} {
	toReturn := iterator.current.data

	return toReturn.(interface{})
}
