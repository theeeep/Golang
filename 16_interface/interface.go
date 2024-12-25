package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment

	fmt.Println("Payment made successfully", amount)
}

func main() {

	payment{}.makePayment(100)
}
