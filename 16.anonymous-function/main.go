package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)

}

// function can also take function as parameter
func transformNumbers(numbers *[]int, transform transformFn) []int {

	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {

	// factor value is locked in inside this scopt that is closure
	return func(number int) int {
		return number * factor
	}
}
