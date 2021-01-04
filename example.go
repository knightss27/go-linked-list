package main

import (
	"fmt"

	"github.com/knightss27/go-linked-list/pkg/list"
)

func main() {
	// Create a new linked list
	myList := list.NewList()

	// Add items
	myList.AddFirst("Item #1")
	myList.AddLast("Item #2")

	// Insert items
	myList.Add(1, "Item #2")
	myList.Set(2, "Item #3")

	// Iterate over items
	iter := myList.Iterator()
	for iter.Next() {
		fmt.Printf("%v, ", iter.Item())
	}

	// Get specific index
	fmt.Printf("\nItem at index 2: %v", myList.Get(2))

	// Remove items
	myList.Remove(1)

	// Clear the list
	myList.Clear()
	fmt.Printf("\nList empty: %v", myList.IsEmpty())

	// Find items in the list
	myList.AddFirst("Item #1")
	fmt.Printf("\nItem #1 in the list: %v, Item #1 has index of: %v", myList.Contains("Item #1"), myList.IndexOf("Item #1"))

	// Get size
	fmt.Printf("\nList current size is: %v", myList.Size())

}
