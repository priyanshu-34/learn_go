package main

import (
	"fmt"
	"time"
)

type order struct {
	id string
	status string
	price float64
	createdAt time.Time
}

func newOrder(status string, price float64) *order{
	myOrder := order{
		id:"123",
		status: status,
		price: price,
		createdAt: time.Now(),
	}
	return &myOrder
}

func (o order) printOrder(){
	fmt.Println("Order details: ", o)
}

func (o *order) updateStatus(newStatus string){
	o.status = newStatus
}

func _struct(){
	fmt.Println("Structs....")

	myOrder := newOrder("pending", 100.0)

	myOrder.printOrder()
	myOrder.updateStatus("Completed")
	myOrder.printOrder()
}