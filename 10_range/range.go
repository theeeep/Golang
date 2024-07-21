package main

import "fmt"

/*
Iteration over data structure
*/
func main() {

	//* Slice
	nums := []int{5, 6, 7}

	// Option: 1 For Loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// Option: 2  Range
	for _, v := range nums {
		fmt.Println(v)
	}

	// Print value with index
	for i, v := range nums {
		fmt.Println("index:", i, "value", v)
	}

	//* Map
	m := map[string]string{"1": "Deep", "2": "Ani", "3": "Ryan"}

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	//* String
	for i, c := range "golang" {
		fmt.Println(i, c) // c hold point rune like some unicode of string
		fmt.Println(i, string(c))
	}
}
