package swordmanboy

import (
	"fmt"
)

func SayTest() {
	fmt.Println("Say Test")

	var fullName string = "Apinun"
	var age = 22

	fmt.Print("Hello, ", fullName, " Your old is ", age)

	// Array slice
	var myArray [3]int
	myArray[0] = 11
	myArray[1] = 12
	myArray[2] = 13

	fmt.Println(myArray)

	myArray[0] = 14

	fmt.Println(myArray)

	for index, item := range myArray {
		fmt.Println("what is ", index, " after", item)
	}

	//Slice
	fmt.Println("MySlice")
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println("mySlice: ", mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	subSlice := mySlice[1:4]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))

	subSlice = append(subSlice, 70)
	fmt.Println(subSlice)

	testStruct()
}

func testStruct() {
	var students [3]Student
	students[0] = Student{
		Name:   "Nun",
		Weight: 81,
		Height: 172,
		Grade:  "S",
	}
	fmt.Println("Student1: ", students[0])
}

type Student struct {
	Name   string
	Weight int
	Height int
	Grade  string
}
