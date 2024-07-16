package main

import (
	"fmt"
	"time"
)

func main() {

	/* Simple switch */

	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	/* Multiple Condition Switch */

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a Weekday :)")
	default:
		fmt.Println("It's Majduri Day :(")
	}

}
