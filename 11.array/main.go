package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:1]
	fmt.Println(prices)
	// if we don't mention any specific size that is slice. otherwise it's an array
	discountPrices := []float64{101.99, 80.99, 20.59}

	// unpacking the array
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}

// func main() {

// 	// array

// 	var productNames [4]string = [4]string{"a", "b"}

// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	// array slice
// 	featuredPrices := prices[1:3]
// 	featuredPrices[0] = 199.99
// 	fmt.Println(featuredPrices)
// 	fmt.Println(prices)
// 	// len represent number of item present in arrary
// 	// cap represent how many items can be stored
// 	fmt.Println(len(productNames), cap(productNames))
// }
