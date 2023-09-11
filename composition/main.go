package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// boats := []*store.Boat{
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Tender", 650.25, 2, true),
	// }

	//lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Boat:", r.Name, "Price:", r.Price(0.2), "Captain:", r.Captain)
	}
}
