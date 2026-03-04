package main

import (
	"fmt"
)

// Struct
type Student struct {
	ID   int
	Name string
	Age  int
}

// Global slice to store students
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
	fmt.Println("Student added successfully!")
}

// View Students
func viewStudents() {
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	fmt.Println("\nStudent List:")
	for _, s := range students {
		fmt.Printf("ID: %d | Name: %s | Age: %d\n", s.ID, s.Name, s.Age)
	}
}

// Search Student
func searchStudent() {
	var id int
	fmt.Print("Enter ID to search: ")
	fmt.Scanln(&id)

	for _, s := range students {
		if s.ID == id {
			fmt.Printf("Found -> ID: %d | Name: %s | Age: %d\n", s.ID, s.Name, s.Age)
			return
		}
	}

	fmt.Println("Student not found.")
}

// Delete Student
func deleteStudent() {
	var id int
	fmt.Print("Enter ID to delete: ")
	fmt.Scanln(&id)

	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Student deleted successfully!")
			return
		}
	}

	fmt.Println("Student not found.")
}

func main() {

	for {
		fmt.Println("\n===== Student Management System =====")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Search Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
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
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
// }


// package main

// import "fmt"

// type Player struct {
//     Name string
//     Age  int
//     City string
// }

// func (p Player) playerDetails() string {
// 	return fmt.Sprintf("{%s %d %s}", p.Name, p.Age, p.City)
// }

// func main() {
//     var player1 Player
// 	// player2 := Player{"Another name", 34, "Another city"}
// 	  //set the value for player1 struct/object property
//     player1.Name = "Any name"
//     player1.Age = 23
//     player1.City = "Any city"
//     // fmt.Println(player1)
// 	// fmt.Println(player2)
// 	fmt.Println(player1.playerDetails())

//     // Output : {Any name 23 Any city}
// }