package main

func calcWithTax(price float64) float64 {
	return price + price*0.2
}

func calcWithoutTax(price float64) float64 {
	return price
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
	insurance.printDetails(0)

	for key, item := range products {
		item.printDetails(key)
	}
}
