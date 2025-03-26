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
	fmt.Print("How much do you want to invest: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Over how many years do you want to invest: ")
	fmt.Scan(&years)
	fmt.Print("What is your expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// Calculations
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Outputs
	fmt.Print("Future value of investment: ")
	fmt.Println(futureValue)
	fmt.Print("Future value of investment adjusted for inflation: ")
	fmt.Println(futureRealValue)
}
