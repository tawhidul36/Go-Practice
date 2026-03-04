package main

import "fmt"

// func Hello() {
//     fmt.Println("Hola!")
// }

// func main() {
//     Hello() //Hello function called and prints “Hola!” in console
// }

// func Repeat(text string, times int) {
//         for i := 0; i < times; i++ {
//                fmt.Println(text)
//        }
//    }

// func main() {
//        Repeat("Hola", 5)  //prints Hola 5 times
//   }


// Varidic function


// func myFunc(numbers ...int) {
//     fmt.Println(numbers)
// }


// func main() {
//     myFunc()
//     myFunc(1)
//     myFunc(1, 2)
//     myFunc(1, 2, 3, 4, 5)
// }

//  func main() {
//      mySlice := []int{1, 2, 3, 4, 5} // initialize slice with elements
//      myFunc("Hello World", mySlice...)
//  }

// func myFunc(mystring string, numbers ...int) {
//      fmt.Println(mystring)
// 	 //print all slice element with range for loop
//      for number := range numbers {
//          fmt.Println(number)
//      }
//  }


//defer function
//print from last to first after main function execution

func main() {
	fmt.Println("Start of main function")
	defer fmt.Println("This will be printed at the end of main function")
	fmt.Println("End of main function")
}