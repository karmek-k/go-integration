package main

import (
	"fmt"
	"strings"

	"github.com/karmek-k/go-integration/functions"
)

func main() {
	fmt.Println("~~~ Numerical Integration ~~~")

	fmt.Println("Available functions:")
	printFunctions(functions.GetDefault())
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
