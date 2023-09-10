package main

import (
	"fmt"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"
	//"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	fmt.Println("Name: ", cart.CustomerName)
	fmt.Println("Total: ", ToCurrency(cart.GetTotal()))
	// color.Red("Name: " + cart.CustomerName)
	// color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
}
