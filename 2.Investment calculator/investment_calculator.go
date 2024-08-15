package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmetAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Invesment Amount: ")
	outputText("Invesment Amount: ")
	fmt.Scan(&investmetAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmetAmount, expectedReturnRate, years)
	// futureValue := investmetAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// outputs
	// fmt.Println("Future Value:", futureValue)

	// fmt.Println("Future Value (Adjusted for Inflation)", futureRealValue)

	// fmt.Printf("Future value: %.2f\nFuture Value (Adjusted for Inflation): %.2f", futureValue, futureRealValue)

	// fmt.Printf(`Future value: %.2f
	// Future Value (Adjusted for Inflation): %.2f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmetAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmetAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
