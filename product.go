package main

import (
	"fmt"
)

type Product struct {
	name, category string
	price          float64
	calcPrice      func(float64) float64
	*Supplier
}

type Supplier struct {
	name, city string
}

type calc func(float64) float64

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	var calcPrice calc
	if price >= 100 {
		calcPrice = calcWithTax
	} else {
		calcPrice = calcWithoutTax
	}
	return &Product{name, category, price, calcPrice, supplier}
}

func (item *Product) printDetails(i int) {
	fmt.Println("Id:", i, "Product:", item.name, "Price:", item.price, "Calc price:", item.calcPrice(item.price),
		"Supplier:", item.Supplier.name)
}
