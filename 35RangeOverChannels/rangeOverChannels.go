package rangeoverchannels

import "fmt"

func Rangeoverchannels() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

//This range iterates over each element as it’s received from queue.
//Because we closed the channel above, the iteration terminates after receiving the 2 elements.
