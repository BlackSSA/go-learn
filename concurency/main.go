package main

import (
	"fmt"
	//"time"
)

func main() {
	// fmt.Println("main function started")
	// CalcStoreTotal(Products)
	// //time.Sleep(time.Second * 5)
	// fmt.Println("main function complete")
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, "for", details.Quantity, details.Product.Name)
	}
	fmt.Println("Channel closed")
}
