package main

import (
	"fmt"
	"os"
	"strings"
)

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

//Function to prompt users
//on whether they would like to end the program
//or go back to the start of the application
func contOrNoT() {
	var yesNo string
	fmt.Println("Would you like to continue from main interface page?")
	fmt.Println("Input Yes to continue, any other input to end program")
	fmt.Scanln(&yesNo)

	if strings.ToLower(yesNo) == "yes" {
		main()
	} else {
		os.Exit(0)
	}
}
