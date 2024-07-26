package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// Passing int
	result := sum(3, 45, 61)
	fmt.Println(result)

	// Passing func
	nums := []int{3, 4, 5}
	fmt.Println(sum(nums...))
}
