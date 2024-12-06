package goroutine

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

func Goroutine() {

	go printMessage("Goroutine 1") // Starts a goroutine
	go printMessage("Goroutine 2") // Starts another goroutine
	go printMessage("Goroutine 3") // Starts another goroutine
	// for i := 0; i < 10; i++ {
	// 	go task(i)
	// }

	time.Sleep(time.Second) // Allow time for goroutines to finish before main exits
}
