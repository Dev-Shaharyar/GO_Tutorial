package timers

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(3 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 is fired")
	}()

	//time.Sleep(2 * time.Second)

	stop := timer2.Stop()

	if stop {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second) //If timer2 had not been stopped, this sleep would have allowed time for "Timer 2 fired" to print.
}
