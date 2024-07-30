package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}
func main() {

	increament := counter()

	fmt.Println(increament())
}
