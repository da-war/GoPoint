package main

import "fmt"

func main() {
	// Code
	age := 32
	hello := &age
	fmt.Println("Age:", age)
	fmt.Println("Hello:", hello)
	fmt.Println("Hello:", *hello)
	fmt.Println("adultYear:", adultYear(age))
}

func adultYear(age int) int {
	return age - 18
}
