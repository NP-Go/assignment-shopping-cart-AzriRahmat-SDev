package main

import "fmt"

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

var categories = []string{"Household", "Food", "Drinks"}
var shoppingList = make(map[string]itemInfo)

func init() {
	shoppingList["Fork"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	shoppingList["Plates"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Cups"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
	shoppingList["Bread"] = itemInfo{category: 1, quantity: 2, unitCost: 2}
	shoppingList["Cake"] = itemInfo{category: 1, quantity: 3, unitCost: 1}
	shoppingList["Coke"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
	shoppingList["Sprite"] = itemInfo{category: 2, quantity: 5, unitCost: 2}
}

func shoppingListDisplay() {
	for key, itemInfo := range shoppingList {
		fmt.Println("Category:", categories[itemInfo.category], "-", "Item:", key, "Quantity:", itemInfo.quantity, "Unit Cost:", itemInfo.unitCost)
	}
}

func calculateByCategory() {
	var totalHHCost float64
	var totalFCost float64
	var totalDCost float64

	for _, itemInfo := range shoppingList {
		if itemInfo.category == 0 {
			totalHHCost += float64(itemInfo.quantity) * itemInfo.unitCost
		} else if itemInfo.category == 1 {
			totalFCost += float64(itemInfo.quantity) * itemInfo.unitCost
		} else {
			totalDCost += float64(itemInfo.quantity) * itemInfo.unitCost
		}

	}
	fmt.Println("Household cost:", totalHHCost)
	fmt.Println("Food cost:", totalFCost)
	fmt.Println("Drinks cost:", totalDCost)
}

func displayByCategory() {
	fmt.Println("List by category")
	for i := 0; i < len(categories); i++ {
		for key, itemInfo := range shoppingList {
			if itemInfo.category == i {
				fmt.Println("Category:", categories[itemInfo.category], "-", "Item:", key, "Quantity:", itemInfo.quantity, "Unit Cost:", itemInfo.unitCost)
			}
		}
	}
}

func addNewItem() {
	for i := 0; i < len(categories); i++ {
		if categories[i] == userInputAddCategory {
			shoppingList[userInputAddItem] = itemInfo{category: i, quantity: userInputAddUnits, unitCost: userInputAddCost}
		}
	}
}

func displayModifiedItem() {
	var itemExist bool = false
	for key, itemInfo := range shoppingList {
		if userInputModifyItemOriginal == key {
			fmt.Println("Current item name is", key, "- Category is", categories[itemInfo.category], "- Quantity is", itemInfo.quantity, "- Unit Cost", itemInfo.unitCost)
			itemExist = true
		}
	}
	if !itemExist {
		fmt.Println("Item enter does not exist")
		modifyItemMenu()
	}
}

func modifyItemsHandler() {
	var tempInfo = shoppingList[userInputModifyItemOriginal]

	if userInputModifyCategory != "" {
		tempInfo = itemInfo{category: userInputModifyCategory}
	} else {
		fmt.Println("No changes to category made.")
	}
	if userInputModifyQty != "" {

	} else {
		fmt.Println("No changes to quantity made.")
	}

	if userInputModifyCost != "" {

	} else {
		fmt.Println("No changes to cost made.")
	}

	if userInputModifyItemNew != "" {

	} else {
		fmt.Println("No changes to name made.")
	}

}
