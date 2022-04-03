package main

type itemData struct {
	Category int
	Quantity int
	unitCost float64
}

func (cs itemData) getCategory() int {
	return cs.Category
}

func (cs itemData) getQuantity() int {
	return cs.Quantity
}

func (cs itemData) getUnitCost() int {
	return int(cs.unitCost)
}
