package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
}

func main() {
	order := order{
		id:        "123",
		amount:    100.0,
		status:    "pending",
		createdAt: time.Now(),
	}

	fmt.Println(order)
	fmt.Println(order.amount)
}
