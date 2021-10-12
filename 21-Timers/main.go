package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(6 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(6 * time.Second)
	fmt.Println(time.Now())
}
