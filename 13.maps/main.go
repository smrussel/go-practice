package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google": "https//google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["AWS"])
	websites["Linkdin"] = "https://linkdin.com"
	fmt.Println(websites)
	delete(websites, "AWS")
	fmt.Println(websites)
}

// structs is used for predefined structure
// struct is used multiple values for same kind or purpose
// maps is like python dictonary key value pair
