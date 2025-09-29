package main

import "fmt"

func main() {
	// Variable declaration
	var name string = "Mohit"
	var age int = 24
	var isStudent bool = false

	// Short variable declaration
	city := "New York"
	height := 5 // or 5.6 will show it as float instead of int
	// height := 5.6

	// Multiple variable declaration
	var country, profession string = "USA", "Engineer"

	// Print variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("City:", city)
	fmt.Println("Height:", height)
	fmt.Println("Country:", country)
	fmt.Println("Profession:", profession)
	fmt.Printf("height type is: %T \n", height)
}