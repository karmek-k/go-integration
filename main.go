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

	format := fmt.Sprintf("|%% %vv|%% %vv|\n", COLUMN_SIZE, COLUMN_SIZE)
	divider := strings.Repeat("-", COLUMN_SIZE)

	fmt.Printf(format, divider, divider)
	fmt.Printf(format, "Name", "Domain")
	fmt.Printf(format, divider, divider)

	for _, f := range funcs {
		fmt.Printf(format, f.Name, f.Domain)
	}

	fmt.Printf(format, divider, divider)
}
