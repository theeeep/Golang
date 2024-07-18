package main

import "fmt"

/* Array: Numbered sequence of specific length */

func main() {
	var nums [4]int // Declaration
	fmt.Println("Length of Array nums:", len(nums))
	nums[0] = 1 // Assign Value
	fmt.Println("Value in  nums[0]:", nums[0])
	fmt.Println("Value in nums[]:", nums)
	/*
		----- O/P -----
		Length of Array nums: 4
		Value in  nums[0]: 1
		Value in nums[]: [1 0 0 0]

	*/

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)
	/*
			----- O/P -----
		[false false true false]
	*/

	/* Single Line Declaration  */
	num := [3]int{1, 2, 3}
	fmt.Println(num)

	/* 2-D Arrays  */
	num2d := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(num2d)
	/*
		----- O/P -----
		 [[3 4] [5 6]]
	*/

	/*

		Cons:
		- Fixed size, that is predictable
		- Memory optimization
		- Constant time access

	*/

}
