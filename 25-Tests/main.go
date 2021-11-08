package main

func MySum(val ...int) int {
	total := 0
	for _, v := range val {
		total += v
	}
	return total
}
