package main

import (
	"fmt"
	"slices"
)

/*
*	Slice -> Dynamic Arrays
* Most used construct in go
* have useful methods
 */
func main() {
	/* Uninitalized slice value contains nil  */ 
	var nums []int
	fmt.Println("Is nil:",nums == nil) // * O/P :- true 

	// var nums2 = make([]int,2) 
	var nums2 = make([]int,0,10) 
	/* 
	make([]int, 0, 10) allocates an underlying array
	of size 10 and returns a slice of length 0 and capacity 10 that is
	backed by this underlying array.
	*/
	fmt.Println("Is nil:",nums2 == nil) 	// * O/P :- false
	fmt.Println("Array values:",nums2) 		// * O/P :- []
	fmt.Println("Capacity",cap(nums2)) 		// * O/P :- 2 (cap -> max number in elements can fit)


	//? Initalizing value to arrays
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	fmt.Println("Array values:",nums2)    // * O/P :-  [1 2 3]
	fmt.Println("Length:",len(nums2)) 		// * O/P :-  5
	fmt.Println("Capacity",cap(nums2)) 		// * O/P :-  8

	//? short-hand 
	nums3 := []int{}
	fmt.Println(nums3)	// * O/P :- []

	var nums4 = make([]int,0,5)
	nums4 = append(nums4, 2)
	var nums5 = make([]int,len(nums4))

	//? Copy function: copy(dest,src)
	copy(nums5,nums4)
	fmt.Println(nums4,nums5) // * O/P :- [2] [2]

	//? Slice Operator

	var nums6 = []int{1,2,3}
	fmt.Println(nums6[0:3])
	/* 
			[0:3] --> 0 include, 3 exclude (0-2)
			short-hand -> [:3] or [1:]
	*/

	//? Slice Equal Method
	var nums7 = []int{1,2,3}
	var nums8 = []int{1,2,3}

	fmt.Println("Slice Methods:",slices.Equal(nums7,nums8))
	/* 
		slice.Equal = two slices are equal: the same length and all elements equal, Equal returns true. 
		If the lengths are different, Equal returns false.
	*/

	//? 2D Dynamic Array
	var nums9 = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(nums9)
} 