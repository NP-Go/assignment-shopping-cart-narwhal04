package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	shopSlice = make([]string, 3)
	mapItem   = make(map[string]itemData)
)

func init() {
	//declaring structs
	cups := itemData{
		Category: 0,
		Quantity: 5,
		unitCost: 3,
	}

	fork := itemData{
		Category: 0,
		Quantity: 4,
		unitCost: 3,
	}

	plates := itemData{
		Category: 0,
		Quantity: 4,
		unitCost: 3,
	}
	bread := itemData{
		Category: 1,
		Quantity: 2,
		unitCost: 2,
	}
	cake := itemData{
		Category: 1,
		Quantity: 3,
		unitCost: 1,
	}
	coke := itemData{
		Category: 2,
		Quantity: 5,
		unitCost: 2,
	}
	sprite := itemData{
		Category: 2,
		Quantity: 5,
		unitCost: 2,
	}

	mapItem = map[string]itemData{ //map declaration with value - of key_value pair as -struct
		"cups":   cups,
		"fork":   fork,
		"plates": plates,
		"cake":   cake,
		"bread":  bread,
		"coke":   coke,
		"sprite": sprite,
	}

	shopSlice = []string{"Household", "Food", "Drinks"} //slice declaration with strings of the various categories

}

func main() {
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report.")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items")
	fmt.Println("5. Delete Item.")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7 Add New Category Name")
	fmt.Println("Select your choice:")

	var inputOne int
	fmt.Println("input")
	fmt.Scan(&inputOne)

	switch inputOne {
	case 1:
		generateList() //generates shopping list with generateList() function
	case 2:
		genReport()
	case 3:
		//this case goes back to main instead of letting code run to conOrNot() function
		addItem()
		// fmt.Println(mapItem)
	case 4:
		modItem()
	default:
		println("You have entered an incorrect entry")
	}
	contOrNoT()

}

//----------------------------------------------------------------------------codes for function-------------------------------------------------------------------------------------------------

func contOrNoT() {
	var yesNo string
	fmt.Println("Would you like to continue from main interface page?")
	fmt.Println("Input Yes to continue, any other input to end program")
	fmt.Scan(&yesNo)

	if strings.ToLower(yesNo) == "yes" {
		main()
	} else {
		os.Exit(0)
	}
}

func generateList() {
	for x, y := range mapItem {
		fmt.Println("category:", shopSlice[y.getCategory()], " - Item:", x, " Quatity:", y.getQuantity(), " Unit Cost:", y.getUnitCost())
	}
}

func genReport() {

	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")

	var inputTwo int
	fmt.Printf("\nChoose your report:\n")
	fmt.Scan(&inputTwo)

	if inputTwo == 1 { //nested for loop to print out total cost by category
		fmt.Println("Total cost by Category.")
		var itemCost int
		for index, category := range shopSlice {
			var catCost int
			for _, value := range mapItem {
				if value.getCategory() == index {
					itemCost = value.getQuantity() * value.getUnitCost()
				} else {
					continue //continue to skip the for loop iterations that do not match for the existing looped category
				}
				catCost += itemCost
			}
			fmt.Println(category, ":", catCost)
		}
	} else if inputTwo == 2 {
		for index, category := range shopSlice {
			for item, value := range mapItem {
				if value.getCategory() == index {
					fmt.Println("category:", category, " - Item:", item, " Quatity:", value.getQuantity(), " Unit Cost:", value.getUnitCost())
				}
			}
		}
	} else if inputTwo == 3 {
		main()
	} else {
		fmt.Println("Program has ended as you have entered an invalid input")
	}

}

func addItem() {
	var (
		nameAddItem string
		catAddItem  string
		unitAddItem int
		costAddItem float64
	)
	fmt.Println("What is the name of your item?")
	fmt.Scan(&nameAddItem)
	fmt.Println("What category does it belong to?")
	fmt.Scan(&catAddItem)
	fmt.Println("How many units are there?")
	fmt.Scan(&unitAddItem)
	fmt.Println("How much does it cost per unit?")
	fmt.Scan(&costAddItem)

	addItemInfo(nameAddItem, catAddItem, unitAddItem, costAddItem)

}

func addItemInfo(nameAddItem string, catAddItem string, unitAddItem int, costAddItem float64) {

	var catAddItemNo int

	if strings.ToLower(catAddItem) == "household" {
		catAddItemNo = 0
	} else if strings.ToLower(catAddItem) == "food" {
		catAddItemNo = 1
	} else if strings.ToLower(catAddItem) == "drinks" {
		catAddItemNo = 2
	} else {
		fmt.Println("Please key in a number from 0 to 2")
		main()
	}

	nameAddNewItem := strings.ToLower(nameAddItem)

	mapItem[nameAddNewItem] = itemData{
		Category: catAddItemNo,
		Quantity: unitAddItem,
		unitCost: costAddItem,
	}

	main()

	fmt.Println(mapItem) //to comment away later - check

}

func modItem() {
	// var (
	// 	itemName string
	// 	// itemNewName string
	// 	// itemNewCat  string
	// 	// itemNewUnit int
	// 	// itemNewCost float32
	// )

	var itemName string

	fmt.Println("Modify Items.")
	fmt.Println("Which item would you wish to modify?")
	fmt.Scan(&itemName)

	// itemVal(itemName)
	fmt.Println(itemName)

	// fmt.Println("Enter new name. Enter for no change.")
	// _, err := fmt.Scan(&itemNewName)
	// fmt.Scanln(&itemNewName)

	// fmt.Print(itemNewName)
	// if itemNewName == "" {
	// 	fmt.Println("error not found")
	// } else {
	// 	fmt.Println("error found")
	// }
	// fmt.Println("Enter new category. Enter for no change.")
	// fmt.Scan(&itemNewCat)
	// fmt.Println("Enter new Quantity. Enter for no change.")
	// fmt.Scan(&itemNewUnit)
	// fmt.Println("Enter new Unit cost. Enter for no change.")
	// fmt.Scan(&itemNewCost)

	// fmt.Printf("%v, %v, %v, %v\n", itemNewName, itemNewCat, itemNewUnit, itemNewCost)

}

func itemVal(itemName string) { //household->food->

	for index, category := range shopSlice {
		for key, value := range mapItem {
			if strings.ToLower(itemName) == key && index == value.getCategory() {
				fmt.Println("Current item name is", key, "- Category is", category, "- Quantity is", value.getQuantity(), "- Unit cost is", value.getUnitCost())
				break
			}
		}
	}
}
