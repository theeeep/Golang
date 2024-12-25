package main

import "fmt"

/* With multiple comment
Now lets say if i want to print string then again i have to create one more func to avoid this we can use generics
*/

func printSlinceGenrics[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// You can also use multiple types as using any is not a good idea
func printSlinceGenrics2[T string | int | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func printSlince(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	printSlince([]int{1, 2, 3})

	// Now here i can use same function to print string as well nums
	printSlinceGenrics([]int{1, 2, 3})
	printSlinceGenrics([]string{"deepak", "behera"})
}
