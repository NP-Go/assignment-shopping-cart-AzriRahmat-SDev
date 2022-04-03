package main

var userInputMainMenu int
var userInputReport int

//Variable for adding
var userInputAddCategory string
var userInputAddItem string
var userInputAddCost float64
var userInputAddUnits int

//Variable for Modify
var userInputModifyItemOriginal string
var userInputModifyItemNew string
var userInputModifyCategory string
var userInputModifyQty int
var userInputModifyCost float64

//Variables for Delete
var userInputDeleteItem string

//Variables for Add new Category
var userInputAddNewCategory string

func main() {

	showMainMenu()

	if userInputMainMenu == 1 {
		shoppingListDisplay()
		main()
	}
	if userInputMainMenu == 2 {
		generateShoppingList()
		if userInputReport == 1 {
			calculateByCategory()
			main()
		} else if userInputReport == 2 {
			displayByCategory()
			main()
		} else {
			main()
		}
	}
	if userInputMainMenu == 3 {
		addItemsMenu()
		main()
	}
	if userInputMainMenu == 4 {
		modifyItemMenu()
		main()
	}

	if userInputMainMenu == 5 {
		deleteItemMenu()
		main()
	}

	if userInputMainMenu == 6 {
		printCurrentDataMenu()
	}

	if userInputMainMenu == 7 {
		printAddNewCategory()
		main()
	}
}
