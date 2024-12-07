package channels

import (
	"fmt"
	"sync"
)

func Channels() {
	// myCh := make(chan int, 3) // buffered channel

	// wg := &sync.WaitGroup{}

	// wg.Add(2)

	// // First goroutine sends multiple values
	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	values := []int{10, 20, 30, 40, 50}
	// 	for _, v := range values {
	// 		fmt.Println("Sending", v, "to the channel")
	// 		ch <- v // Send each value to the channel
	// 		fmt.Println("Sent", v, "to the channel")
	// 	}
	// 	close(ch) // Close the channel to signal that no more values will be sent
	// }(myCh, wg)

	// // Second goroutine reads all the values from the channel
	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	defer wg.Done()
	// 	//for val := range ch { // Read from the channel until it's closed
	// 	fmt.Println("Received from channel:", <-ch)
	// 	//}
	// }(myCh, wg)

	// wg.Wait()
	// fmt.Println("All goroutines finished")
	//OUTPUT
	//only one packet received 10 and all will be in the buffer as buffer size is 4

	ch := make(chan int, 2) // Buffered channel with capacity of 2
	wg := &sync.WaitGroup{}

	// 🔥 Goroutine 1: Producer sending 4 values, but the buffer is only size 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		values := []int{10, 20, 30, 40}
		for _, v := range values {
			fmt.Println("[Producer 1] Sending:", v)
			ch <- v // This will block after sending 2 values because the buffer is full
			fmt.Println("[Producer 1] Sent:", v)
		}
	}()

	// Goroutine 2: Producer sending 3 more values
	wg.Add(1)
	go func() {
		defer wg.Done()
		values := []int{50, 60, 70}
		for _, v := range values {
			fmt.Println("[Producer 2] Sending:", v)
			ch <- v // Will block when channel is full
			fmt.Println("[Producer 2] Sent:", v)
		}
	}()

	// ⚠️ Goroutine 3: Consumer that reads only 3 values (not enough to keep up with producers)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			val := <-ch // This will consume 3 values, but producers send 7 values
			fmt.Println("[Consumer 1] Received:", val)
		}
		fmt.Println("[Consumer 1] Exiting...")
	}()

	// ⚠️ Goroutine 4: Another consumer but only reads 1 value
	wg.Add(1)
	go func() {
		defer wg.Done()
		val := <-ch // Reads only 1 value
		fmt.Println("[Consumer 2] Received:", val)
		fmt.Println("[Consumer 2] Exiting...")
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")

}
