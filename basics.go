package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// Variables
	// basic
	var a int = 5
	b := 7 // this is short hand, as Go identifies the type automatically

	var sum int = a + b
	fmt.Println("this is the sum", sum)

	fmt.Println(" \n------ ARRAYS ------\n ")
	// number of elements in an array are fixed
	var myArray [3]string                            // default way to define an Array
	myArray[1] = "adding value"                      // assigning values
	myArray2 := [3]string{"ma2-1", "ma2-2", "ma2-3"} // short hand way to define an Array

	fmt.Println(myArray)
	fmt.Println(myArray2)

	fmt.Println(" \n------ SLICES ------\n ")
	// Arrays have fixes size so there are 'Slices' they have a dynamic size
	mySlice := []int{4, 2, 6, 7, 10}
	// and later you can add values to it
	mySlice = append(mySlice, 13)

	fmt.Println(mySlice)

	// TODO: Get more slice examples

	fmt.Println(" \n------ MAPS ------\n ")
	// 'Map' is a key value pair in Go
	fruits := make(map[string]int)
	fruits["apple"] = 2
	fruits["banana"] = 3
	fruits["orange"] = 9

	// get all the values
	fmt.Println("all map values -> ", fruits)

	// get the value for a particular key
	fmt.Println("value of a specific key -> ", fruits["orange"])

	// delete something from the map
	delete(fruits, "banana")

	// print the value after delete
	fmt.Println("After deleting a value -> ", fruits)

	fmt.Println(" \n------ PRINT FROM THE LOOP ------\n ")
	// Loops
	// Go has for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(" \n------ FOR AS A WHILE LOOP ------\n ")
	// The for loop also works as a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println(" \n------ ITERATE OVER EACH ELEMENT OF AN ARRAY & MAP ------\n ")
	// for loop can also be used to loop over elements of an Array
	loopingArray := []string{"p", "q", "r"}

	for index, value := range loopingArray {
		fmt.Println("index:", index, "value:", value)
	}

	// similar for maps
	loopingMap := make(map[string]string)
	loopingMap["app"] = "apple"
	loopingMap["goo"] = "google"

	for key, value := range loopingMap {
		fmt.Println("key:", key, "value", value)
	}

	fmt.Println(" \n------ CALLING A FUNCTION ------\n ")
	result := sumOfNumbers(5, 10)
	fmt.Println("Sum of numbers:", result)

	// Now when we call the squareRoot function we will get 2 values
	// We can use if else statement to print the result
	squareResult, error := squareRoot(20) // check the other condition by changing 20 to a negative number

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("squareRoot function result:", squareResult)
	}

	// Using struct
	// The struct is defined below the squareRoot function
	f1 := fruitsType{name: "Apple", color: "Red", count: 4}
	fmt.Println("Value of F1 using struct:", f1)
	fmt.Println("Value of F1's name field:", f1.name)

	// Working with pointer value
	v := 99
	increaseValue(&v) // & is necessary to tell the function that use the value of v at this position
	fmt.Println("Increased pointer value:", v)
}

// Creating a new function
func sumOfNumbers(x int, y int) int {
	return x + y
}

// Little complex function
func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Not for negative values")
	}

	return math.Sqrt(x), nil
}

// A Struct is a collection of fields
type fruitsType struct {
	name  string
	color string
	count int
}

// function returning pointer value
func increaseValue(v *int) { // * tells that use the pointer value insted of a new reference
	fmt.Println(v) // the value of v is the memory reference of the variable
	*v++           // * will make sure the value returned is pointer value not the memory address
}
