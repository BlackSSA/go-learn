package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + price*0.2
}

func calcWithoutTax(price float64) float64 {
	return price
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	acme := &Supplier{"Acme Co", "New York"}
	products := []*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Lifejacket", "Watersports", 48.95, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
		newProduct("Soccer ball", "Soccer", 19.50, acme),
	}

	insurance := newService("Boat Cover", 12, 89.50)
	expensive := []Expense{insurance}
	for _, v := range products {
		expensive = append(expensive, v)
	}

	for _, item := range expensive {
		fmt.Println(item.getName())
	}
}
