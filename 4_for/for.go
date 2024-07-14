package main

import "fmt"

/*
for -> only construct in go for looping
while (Keyword), loop -> ❌
*/
func main() {

	/*
		-- while loop --

			Syntax

		for condition {
		 	do this...
			do that...
		}

	*/

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	/*
		-- infinite loop --

		 	Syntax

		for {
		 	do something
		}

	*/

	/*
		-- CLassic for loop --

		 	Syntax

		for condition{
		 	do something
		}

	*/

	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	/*
		-- Range loop --

		 	Syntax

		for i := range value{
		 	do something
		}

	*/

	for i := range 5 {
		fmt.Println(i)
	}
}
