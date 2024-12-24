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

// Method for struct
// Receiver type is pointer
func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	order := order{
		id:        "123",
		amount:    100.0,
		status:    "pending",
		createdAt: time.Now(),
	}

	fmt.Println("Before change status", order.status)
	order.changeStatus("processing")
	fmt.Println("After change status", order.status)

	fmt.Println(order)
	fmt.Println(order.amount)

	// One line struct
	director := struct {
		firstName     string
		lastName      string
		upcomingMovie string
	}{"Ryan", "Coogler", "Sinners"}

	fmt.Println("Director:", director.firstName, director.lastName, director.upcomingMovie)
}
