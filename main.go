package main

import (
	"fmt"
	"strings"

	"github.com/karmek-k/go-integration/functions"
	"github.com/karmek-k/go-integration/integration"
)

func main() {
	fmt.Println("~~~ Numerical Integration ~~~")

	// TODO: make usable by the user
	execute()
}

func execute() {
	fmt.Println("Available functions:")

	funcs := functions.GetDefault()
	printFunctions(funcs)

	f := funcs[2]
	fmt.Printf("Selected function:\t%v\n", f)

	a := 0.0
	b := 4.0

	fmt.Printf("Integration interval:\t[%v; %v]\n", a, b)

	integral := integration.TrapezoidalIntegral{Function: f, Cuts: 5000}

	result := integral.Calculate(a, b)
	fmt.Printf("Result:\n%.8v\n", result)
}

// printFunctions prints a table containing function names and their domains.
func printFunctions(funcs []functions.Function) {
	const COLUMN_SIZE = 20
	const NUMBER_COLUMN_SIZE = 3

	// e.g. "|% 3v|% 20v|% 20v|\n"
	format := fmt.Sprintf(
		"|%% %vv|%% %vv|%% %vv|\n",
		NUMBER_COLUMN_SIZE,
		COLUMN_SIZE,
		COLUMN_SIZE,
	)
	divider_string := fmt.Sprintf(
		format,
		strings.Repeat("-", NUMBER_COLUMN_SIZE),
		strings.Repeat("-", COLUMN_SIZE),
		strings.Repeat("-", COLUMN_SIZE),
	)

	fmt.Print(divider_string)
	fmt.Printf(format, "#", "Name", "Domain")
	fmt.Print(divider_string)

	for n, f := range funcs {
		fmt.Printf(format, n, f.Name, f.Domain)
	}

	fmt.Print(divider_string)
}
