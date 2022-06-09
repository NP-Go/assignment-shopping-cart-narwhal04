package main

import (
	"fmt"
	"strings"
)

var (
	shopSlice = make([]string, 3)
	mapItem   = make(map[string]itemData)
)

func init() {

	mapItem = map[string]itemData{ //map declaration with value - of key_value pair as -struct
		"cups":   {0, 5, 3},
		"fork":   {0, 4, 3},
		"plates": {0, 4, 3},
		"cake":   {1, 2, 2},
		"bread":  {1, 3, 1},
		"coke":   {2, 5, 2},
		"sprite": {2, 5, 2},
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
	fmt.Scanln(&inputOne)

	switch inputOne {
	case 1:
		generateList() //generates shopping list with generateList() function
	case 2:
		genReport()
	case 3:
		addItem()
	case 4:
		modItem()
	case 5:
		deleteItem()
	case 6:
		printCurData()
	case 7:
		addCat()
	default:
		println("You have entered an incorrect entry")
	}
	contOrNoT()

}

//----------------------------------------------------------------------------codes for function-------------------------------------------------------------------------------------------------

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
	fmt.Scanln(&inputTwo)

	if inputTwo == 1 { //if loop to print out total cost by category
		fmt.Println("Total cost by Category.")
		var itemCost int
		for index, category := range shopSlice { //loop through slice
			var catCost int // cat cost declared here so it will be initiated for every external for loop
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
	} else if inputTwo == 2 { //if loop to print out the struct
		for index, category := range shopSlice {
			for item, value := range mapItem {
				if value.getCategory() == index {
					fmt.Println("category:", category, " - Item:", item, " Quatity:", value.getQuantity(), " Unit Cost:", value.getUnitCost())
				}
			}
		}
	} else if inputTwo == 3 { //return to main
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
	fmt.Scanln(&nameAddItem)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&catAddItem)
	fmt.Println("How many units are there?")
	fmt.Scanln(&unitAddItem)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&costAddItem)

	addItemInfo(nameAddItem, catAddItem, unitAddItem, costAddItem) // function to add add the input to the map

}

func addItemInfo(nameAddItem string, catAddItem string, unitAddItem int, costAddItem float64) {

	catNumber := catConvert(catAddItem)

	nameAddNewItem := strings.ToLower(nameAddItem)

	mapItem[nameAddNewItem] = itemData{ //updates mapItem with new key and value
		Category: catNumber,
		Quantity: unitAddItem,
		unitCost: costAddItem,
	}

	main()

}

// function to covert inputted string category value to int equivalent; 0,1,2
func catConvert(catName string) int {

	var catAddItemNo int
	var catCheckBol bool

	for key, value := range shopSlice {
		if strings.ToLower(catName) == (strings.ToLower(value)) {
			catAddItemNo = key //assign numberic category value to return value
			catCheckBol = false
			break
		} else {
			catCheckBol = true
		}
	}
	if catCheckBol {
		fmt.Println("You have entered the incorrect category input") // this iadded for the incorrect string when converting to index of slice
		contOrNoT()
	}

	return catAddItemNo
}

func modItem() {
	var (
		itemName     string
		itemNewName  string
		itemNewCat   string
		itemNewQuant int
		itemNewCost  float64
		nameBol      bool
	)

	fmt.Println("Modify Items.")
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&itemName)

	itemVal(itemName)

	//-------------------name change section------------------
	fmt.Println("Enter new name. Enter for no change.")

	if _, errName := fmt.Scanln(&itemNewName); errName == nil {
		mapItem[itemNewName] = mapItem[itemName]
		delete(mapItem, itemName)
	} else {
		nameBol = true
		defer fmt.Println("No changes to item name made")
	}

	//-------------------cat change section------------------
	fmt.Println("Enter new category. Enter for no change.")
	_, errCat := fmt.Scanln(&itemNewCat)

	newCat := catConvertMod(itemNewCat)

	if errCat == nil && nameBol == false { //name has been changed and category has been changed
		modCatStruct := mapItem[itemNewName]
		modCatStruct.Category = newCat
		mapItem[itemNewName] = modCatStruct
	} else if errCat == nil && nameBol == true { //category has been changed but name is not changed
		modCatStruct := mapItem[itemName]
		modCatStruct.Category = newCat
		mapItem[itemName] = modCatStruct
	} else {
		defer fmt.Println("No changes to category made")
	}

	//-------------------quantity change section------------------
	fmt.Println("Enter new Quantity. Enter for no change.")
	_, errQuant := fmt.Scanln(&itemNewQuant)

	if errQuant == nil && nameBol == false { //name has been changed and category has been changed
		modQuantStruct := mapItem[itemNewName] //assign updated struct to variable
		modQuantStruct.Quantity = itemNewQuant //assign new quantity to chosen item.quantity
		mapItem[itemNewName] = modQuantStruct  //map updated struct with updated quantity
	} else if errQuant == nil && nameBol == true { //category has been changed but name is not changed
		modQuantStruct := mapItem[itemName]    //assign updated struct to variable
		modQuantStruct.Quantity = itemNewQuant //assign new quantity to chosen item.quantity
		mapItem[itemName] = modQuantStruct     //map updated struct with updated quantity
	} else {
		defer fmt.Println("No changes to quantity made")
	}

	fmt.Println("Enter new Unit cost. Enter for no change.")
	_, errUnitCost := fmt.Scanln(&itemNewCost)

	//-------------------unit cost change section------------------
	if errUnitCost == nil && nameBol == false { //name has been changed and unit cost has been changed
		modUnitCost := mapItem[itemNewName] //assign updated struct to variable
		modUnitCost.unitCost = itemNewCost  //assign new unit cost to chosen item.unitCost
		mapItem[itemNewName] = modUnitCost  //map updated struct with updated unit cost
	} else if errUnitCost == nil && nameBol == true { //category has been changed but name is not changed
		modUnitCost := mapItem[itemName]   //assign updated struct to variable
		modUnitCost.unitCost = itemNewCost //assign new unit cost to chosen item.unitCost
		mapItem[itemName] = modUnitCost    //map updated struct with updated unit cost
	} else {
		defer fmt.Println("No changes to unit cost made")
	}

}

//--------------function to convert string category to int equivalent without return of error message--------
func catConvertMod(catName string) int {

	var catAddItemNo int

	for key, value := range shopSlice {
		if strings.ToLower(catName) == (strings.ToLower(value)) {
			catAddItemNo = key
			break
		}
	}

	return catAddItemNo
}

//--------------------------------------function to print chosen existing item value ----------------------------------------

func itemVal(itemName string) {

	for index, category := range shopSlice {
		for key, value := range mapItem {
			if strings.ToLower(itemName) == key && index == value.getCategory() {
				fmt.Println("Current item name is", key, "- Category is", category, "- Quantity is", value.getQuantity(), "- Unit cost is", value.getUnitCost())
				break
			}
		}
	}
}

//---------------function to delete item from the map : mapItem-----------------
func deleteItem() {

	var (
		delInput  string
		inputBool bool
	)

	fmt.Println("Delete Item.")
	fmt.Println("Enter item name to delete:")
	fmt.Scanln(&delInput)

	for item, _ := range mapItem {
		if strings.ToLower(item) == strings.ToLower(delInput) {
			inputBool = true
			break
		} else {
			continue
		}
	}

	if inputBool {
		fmt.Println("Deleted", delInput)
		delete(mapItem, strings.ToLower(delInput))
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}

//----------------------------------function to print all existing mapItem value--------------------------------------
func printCurData() {
	if len(mapItem) != 0 {
		for key, value := range mapItem {
			fmt.Println(key, "-", value)
		}
	} else {

		fmt.Println("No data found!")
	}
}

//function to add category, check for error input, check for existing input, add new category and update slice-----
func addCat() {
	var newCat string
	var newCatBool bool
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")

	_, err := fmt.Scanln(&newCat)

	for index, category := range shopSlice {
		if err != nil {
			fmt.Println("No Input Found")
			break
		} else if strings.ToLower(newCat) == strings.ToLower(category) {
			fmt.Println("Category:", category, "already exist at index", index)
			newCatBool = false
			break
		} else {
			newCatBool = true
		}
	}

	if newCatBool {
		fmt.Println("New category:", newCat, "added at index", len(shopSlice))
		shopSlice = append(shopSlice, newCat)
	}

}
