package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, -1)
	anotherSum := sumup(1, numbers...)
	fmt.Println(anotherSum)
	fmt.Println(sum)

}

func sumup(startingValue int, numbers ...int) int {

	fmt.Println(startingValue)

	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
