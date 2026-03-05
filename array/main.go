package main

import "fmt"

// -------------- ARRAYS ------------------------

func arrayExample() {

	fmt.Println("--------Array Example---------")

	// Basic array
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Print("Output: ")
	for _, n := range numbers {
		fmt.Print(n, " ")
	}

	fmt.Println()
	fmt.Println()
}

func arrayShortExample() {

	fmt.Println("--------Array Short Declaration---------")

	arr := [3]string{"Go","Python", "C++"}

	fmt.Print("Output: ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println()
}

// -------------- SLICES ------------------------

func sliceBasic() {

	fmt.Println("--------Slice Basic---------")

	s := []int{1, 2, 3, 4}

	fmt.Print("Output: ")
	for _, v := range s {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func sliceAppend() {

	fmt.Println("--------Slice Append---------")

	s := []int{1, 2}

	s = append(s, 3)
	s = append(s, 4, 5)

	fmt.Print("Output: ")
	for _, v := range s {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func sliceFromArray() {

	fmt.Println("--------Slice From Array---------")

	arr := [5]int{10, 20, 30, 40, 50}

	s := arr[1:4]

	fmt.Print("Output: ")
	for _, v := range s {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func sliceLengthCapacity() {

	fmt.Println("--------Slice Length Capacity---------")

	s := []int{10, 20, 30}

	fmt.Println("Output:", "len =", len(s), "cap =", cap(s))

	fmt.Println()
	fmt.Println()
}

// ---------------- MAPS ------------------------

func mapBasic() {

	fmt.Println("--------Map Basic---------")

	m := map[string]int{
		"Rahim": 20,
		"Karim": 22,
	}

	fmt.Print("Output: ")

	for k, v := range m {
		fmt.Print(k, "=", v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func mapCreateAdd() {

	fmt.Println("--------Map Create & Add---------")

	m := make(map[string]string)

	m["language"] = "Go"
	m["database"] = "MySQL"

	fmt.Print("Output: ")

	for k, v := range m {
		fmt.Print(k, ":", v, " ")
	}

	fmt.Println()
	fmt.Println()
}

func mapCheckKey() {

	fmt.Println("--------Map Check Key---------")

	m := map[string]int{
		"id": 100,
	}

	value, ok := m["id"]

	fmt.Println("Output:", value, ok)

	fmt.Println()
	fmt.Println()
}

func mapDelete() {

	fmt.Println("--------Map Delete---------")

	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	delete(m, "a")

	fmt.Print("Output: ")

	for k, v := range m {
		fmt.Print(k, "=", v, " ")
	}

	fmt.Println()
	fmt.Println()
}

// ---------------- MAIN ------------------------

func main() {

	// Arrays
	arrayExample()
	arrayShortExample()

	// Slices
	sliceBasic()
	sliceAppend()
	sliceFromArray()
	sliceLengthCapacity()

	// Maps
	mapBasic()
	mapCreateAdd()
	mapCheckKey()
	mapDelete()

}