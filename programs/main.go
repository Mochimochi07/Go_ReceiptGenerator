package main

import (
	"fmt"
)

type Food struct {
	Name  string
	Price float64
}

var foodList []Food

func main() {
	for {
		fmt.Println("Food List:")
		for _, food := range foodList {
			fmt.Printf("- %s: Php.%.2f\n", food.Name, food.Price)
		}
		fmt.Println("Please choose an option:")
		fmt.Println("1. Add food item")
		fmt.Println("2. Update food item")
		fmt.Println("3. Delete food item")
		fmt.Println("4. Generate receipt")
		fmt.Println("5. Exit")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			addFood()
		case 2:
			updateFood()
		case 3:
			deleteFood()
		case 4:
			generateReceipt()
		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func addFood() {
	var newFood Food
	fmt.Print("Enter food name: ")
	fmt.Scan(&newFood.Name)
	fmt.Print("Enter food price: ")
	fmt.Scan(&newFood.Price)
	foodList = append(foodList, newFood)
}

func updateFood() {
	var name string
	fmt.Print("Enter food name: ")
	fmt.Scan(&name)

	for i, food := range foodList {
		if food.Name == name {
			fmt.Print("Enter new price: ")
			fmt.Scan(&foodList[i].Price)
			return
		}
	}

	fmt.Println("Food not found")
}

func deleteFood() {
	var name string
	fmt.Print("Enter food name: ")
	fmt.Scan(&name)

	for i, food := range foodList {
		if food.Name == name {
			foodList = append(foodList[:i], foodList[i+1:]...)
			return
		}
	}

	fmt.Println("Food not found")
}

func generateReceipt() {
	if len(foodList) == 0 {
		fmt.Println("No food items found")
		return
	}
	total := 0.0
	fmt.Println("Receipt:")
	for _, food := range foodList {
		fmt.Printf("%s: Php.%.2f\n", food.Name, food.Price)
		total += food.Price
	}
	fmt.Printf("Total: Php.%.2f\n", total)
}
