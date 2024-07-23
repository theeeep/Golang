package main

import "fmt"

func main() {
	result := add(2, 5)
	fmt.Println(result)

	fmt.Println(getLanguage())

	// Same as descructive in js/ts
	lang1, lang2, lang3 := getLanguage()
	fmt.Println(lang1, lang2, lang3)

}

// func add(a int, b int) int {
// 	return a + b
// }

/* Another Way of Declaring */
func add(a int, b int) int {
	return a + b
}

func getLanguage() (string, int, bool) {
	return "golang", 2, true
}
