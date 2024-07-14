package main

import "fmt"

func main() {
	const name = "golang"
	// name = "Typescript"  /* Not Allowed */

	/* Multiple constant */
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
