package main

import "fmt"

/////////////////////////////////////////////////
// -------- STRUCT BASICS EXAMPLES -------------
/////////////////////////////////////////////////

// Basic Struct
type Person struct {
	Name string
	Age  int
}

// Nested Struct
type Address struct {
	City    string
	Country string
}

type Employee struct {
	ID      int
	Name    string
	Address Address
}

// Struct Method
func (p Person) greet() {
	fmt.Println("Hello,", p.Name)
}

// Pointer Struct Example
func updateAge(p *Person) {
	p.Age = 30
}

// Struct Basic Demo
func structBasicsDemo() {

	fmt.Println("--------Struct Basic---------")

	// 1. Direct Struct Initialization
	p1 := Person{Name: "Raha", Age: 22}
	fmt.Println("Output:", p1)

	// 2. Field Assignment
	var p2 Person
	p2.Name = "John"
	p2.Age = 25
	fmt.Println("Output:", p2)

	// 3. Pointer Struct
	p3 := &Person{Name: "Alice", Age: 20}
	updateAge(p3)
	fmt.Println("Output:", *p3)

	// 4. Struct Method
	p1.greet()

	// 5. Nested Struct
	emp := Employee{
		ID:   1,
		Name: "Rahim",
		Address: Address{
			City:    "Dhaka",
			Country: "Bangladesh",
		},
	}
	fmt.Println("Output:", emp)

	fmt.Println()
	fmt.Println()
}

/////////////////////////////////////////////////
// -------- STUDENT MANAGEMENT PROJECT ---------
/////////////////////////////////////////////////

// Struct for project
type Student struct {
	ID   int
	Name string
	Age  int
}

// Global slice
var students []Student

// Add Student
func addStudent() {
	var id, age int
	var name string

	fmt.Print("Enter ID: ")
	fmt.Scanln(&id)

	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&age)

	newStudent := Student{
		ID:   id,
		Name: name,
		Age:  age,
	}

	students = append(students, newStudent)

	fmt.Println("Student Added!")
}

// View Students
func viewStudents() {

	fmt.Println("--------Student List---------")

	if len(students) == 0 {
		fmt.Println("No students found")
		return
	}

	for _, s := range students {
		fmt.Printf("ID:%d Name:%s Age:%d\n", s.ID, s.Name, s.Age)
	}
}

// Search Student
func searchStudent() {

	var id int
	fmt.Print("Enter ID: ")
	fmt.Scanln(&id)

	for _, s := range students {
		if s.ID == id {
			fmt.Println("Found:", s)
			return
		}
	}

	fmt.Println("Student not found")
}

// Delete Student
func deleteStudent() {

	var id int
	fmt.Print("Enter ID: ")
	fmt.Scanln(&id)

	for i, s := range students {

		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Student deleted")
			return
		}

	}

	fmt.Println("Student not found")
}

/////////////////////////////////////////////////
// ---------------- MAIN -----------------------
/////////////////////////////////////////////////

func main() {

	// Struct examples first
	structBasicsDemo()

	// Project
	for {

		fmt.Println("===== Student Management =====")
		fmt.Println("1 Add Student")
		fmt.Println("2 View Students")
		fmt.Println("3 Search Student")
		fmt.Println("4 Delete Student")
		fmt.Println("5 Exit")

		var choice int
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			addStudent()

		case 2:
			viewStudents()

		case 3:
			searchStudent()

		case 4:
			deleteStudent()

		case 5:
			fmt.Println("Exit Program")
			return

		default:
			fmt.Println("Invalid choice")
		}

	}

}