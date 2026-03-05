package main

import "fmt"

// ---------- Constants ----------
const country = "Bangladesh"
const maxAge = 100

func main() {

	// ---------- Variables ----------
	name := "Raha"
	age := 22
	a := 10
	b := 20

	// ---------- Operator ----------
	fmt.Println("--------Operator---------")
	sum := a + b
	sub := b - a
	mul := a * b
	div := b / a
	fmt.Println("Output:", sum, sub, mul, div)
	fmt.Println()
	fmt.Println()

	// ---------- Conditional Statement ----------
	fmt.Println("--------Condition---------")
	if age >= 18 {
		fmt.Println("Output:", name, "is adult")
	} else {
		fmt.Println("Output:", name, "is minor")
	}
	fmt.Println()
	fmt.Println()

	// ---------- Switch ----------
	fmt.Println("--------Switch---------")
	day := 5
	switch day {
	case 1:
		fmt.Println("Output: Sunday")
	case 2:
		fmt.Println("Output: Monday")
	case 3:
		fmt.Println("Output: Tuesday")
	default:
		fmt.Println("Output: Another day")
	}
	fmt.Println()
	fmt.Println()

	// ---------- Loop ----------
	fmt.Println("--------Loop---------")
	fmt.Print("Output: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// ---------- Age Loop ----------
	fmt.Println("--------Age Loop---------")
	fmt.Print("Output: ")
	for age < 25 {
		fmt.Print(age, " ")
		age++
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// ---------- Array ----------
	fmt.Println("--------Array---------")
	countries := []string{"Bangladesh", "Japan", "USA"}

	fmt.Print("Output: ")
	for i := 0; i < len(countries); i++ {
		fmt.Print(countries[i], " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// ---------- Range Loop ----------
	fmt.Println("--------Range Loop---------")
	fmt.Print("Output: ")
	for _, c := range countries {
		fmt.Print(c, " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// ---------- Infinite Loop ----------
	fmt.Println("--------Infinite Loop---------")
	count := 1
	fmt.Print("Output: ")
	for {
		fmt.Print(count, " ")
		count++
		if count > 3 {
			break
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// ---------- Final ----------
	fmt.Println("--------Final---------")
	fmt.Println("Output:", name, "from", country, "max age", maxAge)
	fmt.Println()
	fmt.Println()
}