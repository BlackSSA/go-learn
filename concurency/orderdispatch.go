package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrders(channel chan<- DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order Count: ", orderCount)
	for i := 0; i < orderCount; i++ {
		customer := Customers[rand.Intn(len(Customers)-1)]
		product := ProductList[rand.Intn(len(ProductList)-1)]
		quantity := rand.Intn(10)
		channel <- DispatchNotification{customer, product, quantity}
	}
	close(channel)
}
