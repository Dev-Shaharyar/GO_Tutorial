package channelsync

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	fmt.Printf("Worker %d: Starting work\n", id)
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Printf("Worker %d: Done\n", id)
	done <- true // Send signal to main
}

func Channelsync() {
	done := make(chan bool) // Unbuffered channel
	for i := 1; i <= 3; i++ {
		go worker(i, done) // Launch 3 workers
	}

	for i := 1; i <= 3; i++ {
		<-done // Wait for all 3 workers to send a signal
	}

	fmt.Println("All workers completed.")
}
