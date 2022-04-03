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
