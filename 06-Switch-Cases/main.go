package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()

	//Switch-case type 1:
	switch today.Day() {
	case 1:
		fmt.Println("Today is 1th.")
	case 11:
		fmt.Println("Today is 11th.")
	case 15:
		fmt.Println("Today is 15th.")
	case 23:
		fmt.Println("Today is 23th.")
	case 31:
		fmt.Println("Last day.")
	default:
		fmt.Println("No information.")
	}

	//Switch-case type 2:
	switch today.Day() {
	case 1, 3, 5:
		fmt.Println("Go to library.")
	case 10, 20:
		fmt.Println("Go to hospital.")
	case 31:
		fmt.Println("Stay at home.")
	default:
		fmt.Println("No information.")
	}

	//Switch-case type 3:
	switch today.Day() {
	case 9:
		fmt.Println("A")
		fallthrough //:The fallthrough keyword used to force the execution flow to fall through the successive case block
	case 10:
		fmt.Println("B")
	case 11:
		fmt.Println("C")
	case 12:
		fmt.Println("D")
	case 13:
		fmt.Println("E")
	default:
		fmt.Println("XXX")
	}

	//Switch-case type 4:
	switch {
	case today.Day() > 5:
		fmt.Println("Go to library.")
	case today.Day() <= 10:
		fmt.Println("Go to hospital.")
	case today.Day() == 15:
		fmt.Println("Stay at home.")
	default:
		fmt.Println("No information.")
	}

	//Switch-case type 5:
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information.")
	}
}
