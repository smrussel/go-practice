package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	trippled := transformNumbers(&numbers, triple)
	fmt.Println(trippled)
}

// function can also take function as parameter
func transformNumbers(numbers *[]int, transform transformFn) []int {

	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction() transformFn {
	// this function return a function
	return double
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
