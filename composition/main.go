package main

import (
	"composition/store"
	"fmt"
)

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}

	//lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	for _, b := range boats {
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}
}
