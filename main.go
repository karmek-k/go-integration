package main

import (
	"fmt"

	"github.com/karmek-k/go-integration/functions"
	"github.com/karmek-k/go-integration/integration"
	"github.com/karmek-k/go-integration/utils"
)

func main() {
	fmt.Println("~~~ Numerical Integration ~~~")
	fmt.Println()

	// TODO: make usable by the user
	execute()
}

func execute() {
	fmt.Println("Available functions:")

	funcs := functions.Default()
	printFunctions(funcs)

	findex := utils.PromptInRange(0, len(funcs)-1, "Function: ")
	f := funcs[findex]
	a, b := utils.PromptIntegrationInterval()

	fmt.Println("~~~~~~~")
	fmt.Printf("Selected function:\t%v\n", f)
	fmt.Printf("Integration interval:\t[%v; %v]\n", a, b)

	integral := integration.TrapezoidalIntegral{Function: f, Cuts: 5000}

	result := integral.Calculate(a, b)
	fmt.Printf("Result:\n%.5v\n", result)
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
