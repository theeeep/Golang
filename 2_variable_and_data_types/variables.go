package main

import "fmt"

func main() {
	// var name string = "golang"

	var name = "golang" // type infer

	var age int = 30

	fmt.Println(name, age)

	/* short-hand syntax */
	names := "golang" /* --> define and assign at once */
	fmt.Println(names)

	// var price float32 = 50.5
	// var price  = 50.5
	price := 50.5
	fmt.Println(price)
}
