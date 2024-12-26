package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id, "started")
	time.Sleep(time.Second)
	fmt.Println("Task", id, "ended")
}

func main() {
	for i := 1; i <= 10; i++ {
		go task(i)
	}
	time.Sleep(1 * time.Second)
}
