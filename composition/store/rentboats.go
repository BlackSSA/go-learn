package store

type RentalBoat struct {
	*Boat
	IncludeCrew bool
}

func NewRentalBoat(name string, price float64, capacity int, motorized bool, includeCrew bool) *RentalBoat {
	return
}
