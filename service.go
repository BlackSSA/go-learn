package main

import "fmt"

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func newService(description string, durationMonths int, monthlyFee float64) *Service {
	return &Service{description, durationMonths, monthlyFee}
}

func (item *Service) printDetails(i int) {
	fmt.Println("Id:", i, "Service:", item.description, "Duration:", item.durationMonths, "Monthly fee:", item.monthlyFee)
}
