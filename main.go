package main

import (
	"fmt"
	"strconv"

	"github.com/karmek-k/go-integration/functions"
	"github.com/karmek-k/go-integration/integration"
	"github.com/karmek-k/go-integration/utils"
)

func main() {
	fmt.Println("~~~ Numerical Integration ~~~\n")

	// TODO: make usable by the user
	execute()
}

func execute() {
	fmt.Println("Available functions:")

	funcs := functions.Default()
	printFunctions(funcs)

	f := funcs[promptNumberInRange(0, len(funcs)-1)]
	a, b := promptIntegrationInterval()

	fmt.Println("~~~~~~~")
	fmt.Printf("Selected function:\t%v\n", f)
	fmt.Printf("Integration interval:\t[%v; %v]\n", a, b)

	integral := integration.TrapezoidalIntegral{Function: f, Cuts: 5000}

	result := integral.Calculate(a, b)
	fmt.Printf("Result:\n%.8v\n", result)
}

func promptNumberInRange(min, max int) int {
	var number int

	prompt := "Function: "

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

func promptIntegrationInterval() (float64, float64) {
	var numberString, prompt string
	var a, b float64
	var err error

	prompt = "Left integration bound: "

	fmt.Print(prompt)
	fmt.Scan(&numberString)

	a, err = strconv.ParseFloat(numberString, 64)

	for err != nil {
		fmt.Println("Please provide a decimal number.")

		fmt.Print(prompt)
		fmt.Scan(&numberString)

		a, err = strconv.ParseFloat(numberString, 64)
	}

	prompt = "Right integration bound: "

	fmt.Print(prompt)
	fmt.Scan(&numberString)

	b, err = strconv.ParseFloat(numberString, 64)

	for err != nil {
		fmt.Println("Please provide a decimal number.")

		fmt.Print(prompt)
		fmt.Scan(&numberString)

		b, err = strconv.ParseFloat(numberString, 64)
	}

	return a, b
}

// printFunctions prints a table containing function names and their domains.
func printFunctions(funcs []functions.Function) {
	const COLUMN_SIZE = 20
	const NUMBER_COLUMN_SIZE = 3

	table := utils.NewTable(
		NUMBER_COLUMN_SIZE,
		COLUMN_SIZE,
		COLUMN_SIZE,
	)

	fmt.Println(table.Separator())
	fmt.Println(table.Values("#", "Name", "Domain"))
	fmt.Println(table.Separator())

	for n, f := range funcs {
		fmt.Println(table.Values(n, f.Name, f.Domain))
	}

	fmt.Println(table.Separator())
}
