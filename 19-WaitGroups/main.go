package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second * 3)
	fmt.Printf("Worker %d done\n", id)
}

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("Hello World")
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
	fmt.Println("Finished Execution")
}
