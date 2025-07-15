package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	phone string
}

//order struct
type order struct {
	id int
	amount float32
	status string
	customer customer
	createdAt time.Time //precision: nanosecond 
}


func  newOrder(id int, amount float32, status string) *order{
	//initial function or constructor
		myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

//receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	myOrder := newOrder(111, 30.30, "pending")
	fmt.Println(myOrder.amount)

	newOrder := order{id: 1,amount: 299,status: "True",customer: customer{name: "Ayush",phone: "9813493440"},createdAt: time.Now()}

	language := struct{
		name string
		isGood bool
	}{
		"English",
		true,
	}
	newOrder.customer.name = "Duck"
	fmt.Println(language)
	fmt.Println(newOrder)

}

