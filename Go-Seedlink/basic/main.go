package main

import (
	"fmt"
)


func main() {

	var array [5]int
	fmt.Print("Array length: ", len(array), "\n")

	//Array
	// var array [5]int
	// for i:=0;i<5;i++{
	// 	fmt.Printf("Enter element %d: ", i+1)
	// 	fmt.Scanln(&array[i])
	// }
	// for i:=0;i<5;i++{

	// 	fmt.Println(array[i])
	// }

	//Slice
	// var slice []int
	// var n int
	// fmt.Print("Enter the number of elements: ")
	// fmt.Scanln(&n)
	// for i:=0;i<n;i++{
	// 	var a int
	// 	fmt.Printf("Enter element %d: ", i+1)
	// 	fmt.Scanln(&a)
	// 	slice = append(slice, a)
	// }
	// for i:=0;i<len(slice);i++{
	// 	fmt.Println(slice[i])
	// }
	// slice = append(slice, 6)
	// fmt.Println("After appending 6: ", slice)


	//Map
	// var mp map[int]string
	// mp = make(map[int]string)

	// var n int
	// fmt.Print("Enter the number of elements: ")
	// fmt.Scanln(&n)
	// for i:=0;i<n;i++{
	// 	var key int 
	// 	var value string 
	// 	fmt.Printf("Enter key %d: ", i+1)
	// 	fmt.Scanln(&key)
	// 	fmt.Printf("Enter value for key %d: ", key)
	// 	fmt.Scanln(&value)
	// 	mp[key] = value
	// }

	// for key,value:=range mp{
	// 	fmt.Printf("Key: %d, Value: %s\n", key, value)
	// }



	//Conditional Statements
	// var a int
	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&a)
	// if(a>=18){
	// 	fmt.Println("You are an adult")
	// }else{
	// 	fmt.Println("You are a minor")
	// }


	//Loops
	// for i:=1;i<=3;i++{
	// 	var a int
	// 	fmt.Printf("Enter number %d: ", i)
	// 	fmt.Scanln(&a)
	// }
	// for i:=1;i<=3;i++{
	// 	fmt.Printf("Number %d: %d\n", i, a)
	// }



}