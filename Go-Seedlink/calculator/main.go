package main

import (
	"fmt"
)

func mixed(a,b float64) (float64, float64, float64, float64) {
	return a + b, a - b, a * b, divide(a, b)
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Cannot divide by zero")
		return 0
	}
	return a / b
}

func main() {
	var a, b float64
	var choice int

	for { // infinite loop

		fmt.Println("\nSimple Calculator")
		fmt.Println("-----------------")
		fmt.Println("Press 1 for addition")
		fmt.Println("Press 2 for subtraction")
		fmt.Println("Press 3 for multiplication")
		fmt.Println("Press 4 for division")
		fmt.Println("Press 5 for all operations")
		fmt.Println("Press any other number to Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// Break condition
		if choice < 1 || choice > 5 {
			fmt.Println("Exit calculator...")
			break
		}

		fmt.Print("Enter the first number: ")
		fmt.Scanln(&a)

		fmt.Print("Enter the second number: ")
		fmt.Scanln(&b)

		switch choice {
		case 1:
			fmt.Printf("The sum of %.2f and %.2f is %.2f\n", a, b, add(a, b))
		case 2:
			fmt.Printf("The difference of %.2f and %.2f is %.2f\n", a, b, subtract(a, b))
		case 3:
			fmt.Printf("The product of %.2f and %.2f is %.2f\n", a, b, multiply(a, b))
		case 4:
			fmt.Printf("The quotient of %.2f and %.2f is %.2f\n", a, b, divide(a, b))
		
		case 5:
			sum, diff, prod, quot := mixed(a, b)
			fmt.Printf("For %.2f and %.2f:\n", a, b)
			fmt.Printf("  Sum: %.2f\n", sum)
			fmt.Printf("  Difference: %.2f\n", diff)
			fmt.Printf("  Product: %.2f\n", prod)
			fmt.Printf("  Quotient: %.2f\n", quot)
		}
	}
}



