package main

import (
	"fmt"
)

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 275},
	{"Lifejacket", "Watersports", 48.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func main() {
	fmt.Println("Hello, World!")
	dict := make(map[string][]*Product)
	dict["Watersports"] = make([]*Product, 0, 2)
	dict["Watersports"] = append(dict["Watersports"], ProductList[0])
	dict["Watersports"] = append(dict["Watersports"], ProductList[1])
	fmt.Println(&ProductList[0], *ProductList[0], ProductList[0])
	fmt.Println(&(dict["Watersports"][0]), dict["Watersports"][0])
	ProductList[0].Price = 100
	fmt.Println(&ProductList[0], *ProductList[0], ProductList[0])
	fmt.Println(&(dict["Watersports"][0]), dict["Watersports"][0])
}
