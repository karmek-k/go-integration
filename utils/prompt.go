package utils

import (
	"fmt"
	"strconv"
)

// PromptInRange prompts for an integer in a given range.
func PromptInRange(min, max int, prompt string) int {
	var number int

	fmt.Print(prompt)
	fmt.Scan(&number)

	// if out of bonds
	for number < min || number > max {
		fmt.Printf(
			"Provided number should not be lesser than %v and greater than %v.\n",
			min, max)

		fmt.Print(prompt)
		fmt.Scan(&number)
	}

	return number
}

// PromptIntegrationInterval prompts for the bounds of a definite integral.
func PromptIntegrationInterval() (float64, float64) {
	a := promptFloat("Left integration bound: ")
	b := promptFloat("Right integration bound: ")

	return a, b
}

// promptFloat prompts for a floating point number.
func promptFloat(prompt string) float64 {
	var numberString string
	var result float64
	var err error

	fmt.Print(prompt)
	fmt.Scan(&numberString)

	result, err = strconv.ParseFloat(numberString, 64)

	for err != nil {
		fmt.Println("Please provide a decimal number.")

		fmt.Print(prompt)
		fmt.Scan(&numberString)

		result, err = strconv.ParseFloat(numberString, 64)
	}

	return result
}
