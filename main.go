package main

import (
	"fmt"
	"math"
)

func main() {
	// Variables
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	// Variable inputs
	fmt.Print("How much do you want to invest: R")
	fmt.Scan(&investmentAmount)
	fmt.Print("Over how many years do you want to invest: ")
	fmt.Scan(&years)
	fmt.Print("What is your expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// Calculations
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Outputs
	fmt.Printf("Future value of investment: R%.2f\n", futureValue)
	fmt.Printf("Future value of investment adjusted for inflation: R%.2f\n", futureRealValue)
}
