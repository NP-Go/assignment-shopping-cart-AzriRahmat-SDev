package main

var userInputMainMenu int
var userInputReport int
var userInputAddCategory string
var userInputAddItem string
var userInputAddCost float64
var userInputAddUnits int

var userInputModifyItemOriginal string
var userInputModifyItemNew string
var userInputModifyCategory string
var userInputModifyQty int
var userInputModifyCost float64

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
}
