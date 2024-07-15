package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("Person is Adult")
	} else if age >= 12 {
		fmt.Println("Person is Teenager")
	} else {
		fmt.Println("Person is Minor")
	}

	/*	We can declare a variable inside if construct */

	if age := 15; age >= 18 {

		fmt.Println("You are 18+")
	} else {
		fmt.Println("Go and and watch shinchan broo 😒")
	}

	role := "Admin"
	hasPermission := true

	if role == "Admin" && hasPermission == true {
		fmt.Println("Admin access grandted!")
	} else {
		fmt.Println("Admin access not granted!")
	}

	/* Go doesn't have ternery (: ?) operator */
}
