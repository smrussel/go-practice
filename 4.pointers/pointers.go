package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int

	agePointer := &age

	fmt.Println("AgePointer:", agePointer)
	fmt.Println("AgePointer value:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
