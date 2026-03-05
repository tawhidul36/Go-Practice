package main

import "fmt"

// -------- Simple Function --------
func Hello() {
	fmt.Println("Hello World!")
}

// -------- Function with Parameters --------
func Repeat(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(text, " ")
	}
	fmt.Println()
}

// -------- Function Returning Value --------
func Sum(a int, b int) int {
	return a + b
}

// -------- Variadic Function --------
func PrintNumbers(numbers ...int) {
	for _, n := range numbers {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

// -------- Function with Slice + Variadic --------
func PrintMessageNumbers(msg string, numbers ...int) {
	fmt.Print(msg, ": ")
	for _, n := range numbers {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

// -------- Defer Example Function --------
func DeferExample() {
	fmt.Println("Inside function")
	defer fmt.Println("Deferred line (runs last)")
	fmt.Println("Before defer executes")
}

// -------- Main Function --------
func main() {

	fmt.Println("--------Simple Function---------")
	fmt.Print("Output: ")
	Hello()
	fmt.Println()
	fmt.Println()

	fmt.Println("--------Function With Parameters---------")
	fmt.Print("Output: ")
	Repeat("Hello", 5)
	fmt.Println()
	fmt.Println()

	fmt.Println("--------Return Function---------")
	fmt.Print("Output: ")
	result := Sum(10, 20)
	fmt.Println(result)
	fmt.Println()
	fmt.Println()

	fmt.Println("--------Variadic Function---------")
	fmt.Print("Output: ")
	PrintNumbers(1, 2, 3, 4, 5)
	fmt.Println()
	fmt.Println()

	fmt.Println("--------Slice to Variadic Function---------")
	mySlice := []int{10, 20, 30}
	fmt.Print("Output: ")
	PrintMessageNumbers("Numbers", mySlice...)
	fmt.Println()
	fmt.Println()

	fmt.Println("--------Defer Function---------")
	fmt.Print("Output:\n")
	DeferExample()
	fmt.Println()
	fmt.Println()
}