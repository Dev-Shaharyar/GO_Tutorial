package thicker

import (
	"fmt"
	"time"
)

func checkAPI() {
	fmt.Println("Checking API for new messages at", time.Now())
}
func Thicker() {

	// ticker := time.NewTicker(time.Second)
	// //done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case t := <-ticker.C:
	// 			fmt.Println("Tick at", t)
	// 		}
	// 	}
	// }()

	// time.Sleep(5 * time.Second)
	// ticker.Stop()
	// // done <- true
	// //fmt.Println("Ticker stopped")

	ticker := time.NewTicker(5 * time.Second) // Call API every 5 seconds

	go func() {
		for range ticker.C {
			checkAPI() // Call the API
		}
	}()

	// Run for 30 seconds, then stop the ticker
	time.Sleep(30 * time.Second)
	ticker.Stop()
	fmt.Println("Stopped checking API")
}

//OUTPUT
//print checking API...... at... -> 5 times
