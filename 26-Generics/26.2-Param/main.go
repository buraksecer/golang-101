package main

import (
	"fmt"
	"time"
)

type Param interface {
	[]int | []string | time.Time | string
}

func Operation[T Param](list T) {
	fmt.Println("Param", list)
}

func main() {
	Operation("selam")
	Operation([]string{"burak", "selam"})
	Operation(time.Now())
}
