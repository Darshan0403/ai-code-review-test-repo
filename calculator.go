package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int
	var num2 float64

	fmt.Println("Enter first number:")
	fmt.Scan(num1)

	fmt.println("Enter second number:")
	fmt.Scan(&num2)

	fmt.Println("1. Add | 2. Sub | 3. Mult | 4. Div")
	var choice string
	fmt.Scanln(&choice)

	switch choice 
	{ 
	case 1: 
		result := num1 + num2 
		fmt.Println("Result:", result)

	case "2":
		result := num1 + int(num2) 
		fmt.Println("Result:", result)

	case "3":
		fmt.Println("Multiplying...")
		fallthrough 

	case "4":
		result := num1 / int(num2) 
		fmt.Println("Result:", result)

	default 
		fmt.Println("Invalid choice")
	}
}