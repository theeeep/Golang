package main

import "fmt"

/*
----- By Value -----
When you pass a variable "by value," you are giving a copy of that variable.
This means that any changes made to the copy do not affect the original variable
*/
func changeNum(num int) {
	num = 7
	fmt.Println("In changeNum:", num)
}

/*
----- By Reference -----
When you pass a variable "by reference," you are giving a pointer to that variable.
This means that any changes made to the variable through the pointer affect the original variable
*/
func changeNumPointer(num *int) {
	*num = 7
	fmt.Println("In changeNumPointer:", *num)
}

/*
----- Multiple Pointers -----
When you pass multiple variables to a function, each variable is a pointer to the original variable
*/
func changeNumPointers(num1 *int, num2 *int) {
	*num1 = 7
	*num2 = 8
	fmt.Println("In changeNumPointers:", *num1, *num2)
}

func main() {
	num := 3
	num2 := 4

	changeNum(num)

	changeNumPointer(&num)

	fmt.Println("Num change in changeNum func", num)
	fmt.Println("Num in changeNumPointer", &num)

	changeNumPointers(&num, &num2)

	fmt.Println("Num change in changeNumPointers func", num)
	fmt.Println("Num in changeNumPointers", &num)
	fmt.Println("Num change in changeNumPointers func", num2)
	fmt.Println("Num in changeNumPointers", &num2)
}
