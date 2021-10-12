package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan int)
	go func() { message <- 31 }()
	test := <-message
	fmt.Println(test)

	messageString := make(chan string, 3)
	messageString <- "1."
	messageString <- "2."
	messageString <- "3."
	fmt.Println(<-messageString)
	fmt.Println(<-messageString)
	fmt.Println(<-messageString)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	waitVal := make(chan int, 1)
	val := make(chan int, 1)
	waitValue(waitVal, 3)
	resultValue(waitVal, val)
	fmt.Println(<-val)

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func worker(done chan bool) {
	fmt.Println("starting...")
	time.Sleep(time.Second)
	fmt.Println("done!")
	done <- true
}

func waitValue(value chan<- int, val int) {
	value <- val * 3
}

func resultValue(value <-chan int, rVal chan<- int) {
	v := <-value
	rVal <- v
}