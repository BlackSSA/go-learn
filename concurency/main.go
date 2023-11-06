package main

import (
	"fmt"
	//"time"
)

func reciveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, "for", details.Quantity, details.Product.Name)
	}
	fmt.Println("Channel closed")
}

func main() {
	// fmt.Println("main function started")
	// CalcStoreTotal(Products)
	// //time.Sleep(time.Second * 5)
	// fmt.Println("main function complete")
	dispatchChannel := make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
	reciveDispatches((<-chan DispatchNotification)(dispatchChannel))
}
