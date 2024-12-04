package maps

import (
	"fmt"
	"maps"
)

func Maps() {
	//using make
	m := make(map[string]int)

	// Add elements to the map
	m["apple"] = 5
	m["banana"] = 3

	// Update an element
	m["apple"] = 10

	// Access an element
	value, ok := m["apple"]
	if ok {
		fmt.Println("apple:", value)
	} else {
		fmt.Println("apple not found")
	}

	// Check for a key that doesn't exist
	value, ok = m["orange"]
	if ok {
		fmt.Println("orange:", value)
	} else {
		fmt.Println("orange not found")
	}

	fmt.Println("Map contents:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	// Delete a key
	delete(m, "banana")

	// clear(m)
	// fmt.Println("map:", m)   //will clear all key value pairs

	// Verify deletion
	fmt.Println("After deletion:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 3}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
