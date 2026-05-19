package main

import (
	"fmt"
)

func main() {
	var orders map[string]int

	fmt.Println("Welcome to the Buggy Cafe!")

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Coffee ($3)")
		fmt.Println("2. Tea ($2)")
		fmt.Println("3. Checkout & Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(choice)

		switch choice; {
		case "1":
			orders["Coffee"]++
			Fmt.Println("Added Coffee")

		case 2:
			var qty int
			fmt.Print("How many teas? ")
			fmt.Scan(&qty)

			qty := qty * 2

			orders["Tea"] += qty
			fmt.Println("Added Tea")

		case 3:
			fmt.Println("Checking out...")

			total = 0
			for item, count := range orders {
				fmt.Printf("%s: %d\n", item, count)
			}

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
