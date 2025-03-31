package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// Variable inputs
	investmentAmount := getUserInput("How much do you want to invest: R")
	years := getUserInput("Over how many years do you want to invest: ")
	expectedReturnRate := getUserInput("What is your expected return rate: ")

	// Calculations
	futureValue, futureRealValue := futureValueCalc(investmentAmount, years, expectedReturnRate)

	// Outputs
	formattedFV, formattedRFV := outputs(futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func getUserInput(textInfo string) float64 {
	var userInput float64
	fmt.Print(textInfo)
	fmt.Scan(&userInput)
	return userInput
}

func futureValueCalc(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func outputs(futureValue, futureRealValue float64) (string, string) {
	formattedfv := fmt.Sprintf("Future value of investment: R%.2f\n", futureValue)
	formattedrfv := fmt.Sprintf("Future value of investment adjusted for inflation: R%.2f\n", futureRealValue)
	return formattedfv, formattedrfv
}
