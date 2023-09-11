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

	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
	// 	store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	// }

	// product := store.NewProduct("Kayak", "Watersports", 279)
	// deal := store.NewSpecialDeal("Weekend Special", product, 50)

	// Name, price, Price := deal.GetDetails()

	// for _, r := range rentals {
	// 	fmt.Println("Boat:", r.Name, "Price:", r.Price(0.2), "Captain:", r.Captain)
	// }

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

	//fmt.Println("Name:", Name, "Price field:", price, "Price method:", Price)
	//446
}
