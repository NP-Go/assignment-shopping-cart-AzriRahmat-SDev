package main

import "fmt"

func showMainMenu() {
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items.")
	fmt.Println("5. Delete Items.")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7. Add New Category Name")
	fmt.Println("Select your choice: ")
	fmt.Scanln(&userInputMainMenu)
}

func generateShoppingList() {
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")
	fmt.Println("Choose your report: ")
	fmt.Scanln(&userInputReport)
}

func addItemsMenu() {
	fmt.Println("What is your item?")
	fmt.Scanln(&userInputAddItem)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&userInputAddCategory)
	fmt.Println("How many unit are there?")
	fmt.Scanln(&userInputAddUnits)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&userInputAddCost)
	addNewItem()
}

func modifyItemMenu() {
	fmt.Println("Modify Item.")
	fmt.Println("Which Item would you wish to modify")
	fmt.Scanln(&userInputModifyItemOriginal)
	displayModifiedItem()
	fmt.Println("Enter new name. Enter for no change")
	fmt.Scanln(&userInputModifyItemNew)
	fmt.Println("Enter new Category. Enter for no change")
	fmt.Scanln(&userInputModifyCategory)
	fmt.Println("Enter new Quantity. Enter for no change")
	fmt.Scanln(&userInputModifyQty)
	fmt.Println("Enter new Unit Cost. Enter for no change")
	fmt.Scanln(&userInputModifyCost)
	modifyItemsHandler()
}

func deleteItemMenu() {
	fmt.Println("Delete Item.")
	fmt.Println("Enter item name to delete: ")
	fmt.Scanln(&userInputDeleteItem)
	deleteItemHandler()
}

func printCurrentDataMenu() {
	fmt.Println("Print Current Data.")
	printDataHandler()
}
