package main

import "fmt"

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// make slice with make function
	userNames := make([]string, 2, 5)

	userNames[0] = "julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	courseRatings.output()
	// fmt.Println(courseRatings)

	// for loop for -array or slice
	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	// for loop for -maps
	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}

}
