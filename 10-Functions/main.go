package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {

	BasicHello()

	Total(10, 3)

	fmt.Println("Total:", TotalReturn(1, 3))

	a, b := MultiReturn()

	fmt.Println(a, b)

	var number = 20
	var text = "John"
	fmt.Println("Before:", text, number)
	update(&number, &text)
	fmt.Println("After :", text, number)

	area(1, 2)

	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	partial := partialSum(3)
	fmt.Println(partial(7))

}

func BasicHello() {
	fmt.Println("Hello!")
}

func Total(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func TotalReturn(x int, y int) int {
	return x + y
}

func MultiReturn() (int, string) {
	return 1, "YES"
}

func update(a *int, t *string) {
	*a = *a + 5
	*t = *t + " Doe"
	return
}

func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
