package main

import (
	"fmt"
	"maps"
)

/*
Maps --> Key:Value
Syntax :- map[key Dtype] value Dtype
ex:- m := make(map[string]string)

if a map is empty and its value data type value will be
if string -> empty value, int --> 0, boolean --> false
*/
func main() {
	// Declaration Map
	m := make(map[string]string)

	// Assign element
	m["name"] = "golang"
	m["age"] = "21"

	// Return element
	fmt.Println(m["name"], m["age"])
	fmt.Println(len(m))

	//Remove element
	delete(m, "age")
	fmt.Println(len(m))

	// to clear map
	clear(m)

	// short-hand
	n := map[string]int{"price": 40, "phones": 3, "sell": 4}
	fmt.Println(n)

	v, has := n["phones"] // v --> hold value of key (v = 3),  has --> hold boolean value, key exist or not in map (has = true)
	fmt.Println(v, has)

	// Check two maps are equal or not with its key:value
	m1 := map[string]string{"name": "Deep", "age": "22", "cinephile": "true"}
	m2 := map[string]string{"name": "Ryan", "age": "28", "cinephile": "true"}

	fmt.Println(maps.Equal(m1, m2))
}
