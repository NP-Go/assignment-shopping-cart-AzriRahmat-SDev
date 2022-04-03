package main

import "fmt"

type itemInfo struct {
	category int
	quantity int
	unitCost float64
}

// type empty struct {
// }

var categories = []string{"Household", "Food", "Drinks"}
var shoppingList = make(map[string]itemInfo)

// var shoppingListEmpty = make(map[string]empty)

func init() {
	shoppingList["Fork"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	shoppingList["Plate"] = itemInfo{category: 0, quantity: 4, unitCost: 3}
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
	// shoppingList["Fork"] = itemInfo{category: 0, quantity: 5, unitCost: 3}
	// User enters "Fork"

	var tempInfoForItem = shoppingList[userInputModifyItemOriginal]
	// tempInfoForItem = shoppingList["Fork"]
	// tempInfoForItem = itemInfo{category: 0, quantity: 5, unitCost: 3}

	// tempInfoForItem {
	// 	  category: 0,
	//    quantity: 5,
	//	  cost: 3
	// }

	var setUserInputModifyCategory int

	for i := 0; i < len(categories); i++ {
		if categories[i] == userInputModifyCategory {
			setUserInputModifyCategory = i
		}
	}

	// User input category "Food"
	// userInputModifyCategory = 1

	if userInputModifyCategory != "" {
		tempInfoForItem.category = setUserInputModifyCategory

		// tempInfoForItem {
		// 	  category: 1,
		//    quantity: 5,
		//	  cost: 3
		// }

	} else {
		fmt.Println("No changes to category made.")
	}

	if userInputModifyQty != 0 {
		tempInfoForItem.quantity = userInputModifyQty
	} else {
		fmt.Println("No changes to quantity made.")
	}

	if userInputModifyCost != 0 {
		tempInfoForItem.unitCost = userInputModifyCost
	} else {
		fmt.Println("No changes to cost made.")
	}

	// tempInfoForItem {
	// 	  category: 1,
	//    quantity: 45,
	//	  cost:100
	// }

	//Scenario 1: User input blank
	//  - Update "Fork" with new tempInfoForItem
	if userInputModifyItemNew == "" {

		// i, exist := shoppingList["Fork"]
		// i => itemInfo{category:0, quantity:1, cost:10}
		// exist => true

		// i, exist := shoppingList["mouse"]
		// i => 0
		// exist => false
		shoppingList[userInputModifyItemOriginal] = tempInfoForItem
		fmt.Println("No changes to cost made.")
	}
	//Scenario 2: User input new item name
	//  - Add "Plate" with new tempInfoForItem
	//  - Delete "Fork"
	if userInputModifyItemNew != "" {
		decision := "Yes"
		_, exist := shoppingList[userInputModifyItemNew]
		fmt.Println("testing", exist)
		if exist {
			fmt.Println("The item you keyed already exist, are you sure you want to overide? Yes or No")
			fmt.Scanln(&decision)
		}
		if decision == "Yes" {
			shoppingList[userInputModifyItemNew] = tempInfoForItem
			delete(shoppingList, userInputModifyItemOriginal)
		} else {
			fmt.Println("No change made RTB")
		}
	}
}

func deleteItemHandler() {
	for key := range shoppingList {
		if userInputDeleteItem == key {
			delete(shoppingList, key)
		}
	}
	_, exist := shoppingList[userInputDeleteItem]
	if !exist {
		fmt.Println("Item not found. nothing to delete!")
	}
}

func printDataHandler() {
	for key, itemInfo := range shoppingList {
		fmt.Println(key, itemInfo)
	}
	if len(shoppingList) <= 0 {
		fmt.Println("No data found")
	}
}

func AddNewCategoryHandler() {
	// categories = {"Household", "Food", "Drinks"}

	//Check through categories to see if it exists already
	//Reject if exists

	//If not exists, add to categories
	var exist bool = false
	var currentIndex int
	for i := 0; i < len(categories); i++ {
		// fmt.Println(userInputAddNewCategory, "==", categories[i])
		// fmt.Println(userInputAddNewCategory == categories[i])
		if userInputAddNewCategory == categories[i] {
			exist = true
			currentIndex = i
		}
	}

	if exist {
		fmt.Println("Category", userInputAddNewCategory, "already exist at index", currentIndex)
	} else {
		categories = append(categories, userInputAddNewCategory)
		fmt.Println("New Category:", userInputAddNewCategory, "added at index", len(categories)-1)
		fmt.Println(categories)
	}

}
