package main

import (
	"fmt"
)

// recursion - the function call itself

func main() {
	fact := factorial(3)
	fmt.Println(fact)
}

// factorial of 5 => 5*4*3*2*1

// normal approach

// func factorial(number int) int {
// 	result := 1
// 	for i := 1; i <= number; i++ {
// 		result = result * i
// 	}

// 	return result
// }

// recursion approach

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
