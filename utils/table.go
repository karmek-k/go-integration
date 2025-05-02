package utils

import (
	"fmt"
	"strings"
)

type Table struct {
	format    string
	separator string
}

// NewTable creates an ASCII table with a default format.
func NewTable(sizes ...int) Table {
	return Table{
		format:    makeFormat(&sizes),
		separator: makeSeparator(&sizes),
	}
}

// Values gives a formatted table row with values.
func (t *Table) Values(values ...any) string {
	return fmt.Sprintf(t.format, values...)
}

// Separator gives a separator table row.
func (t *Table) Separator() string {
	return t.separator
}

func makeFormat(sizes *[]int) string {
	// something like "|%% %vv|%% %vv|%% %vv|\n"
	format := "|"

	for _, size := range *sizes {
		format += fmt.Sprintf("%% %vv|", size)
	}

	return format
}

func makeSeparator(sizes *[]int) string {
	separator := "|"

	for _, size := range *sizes {
		separator += fmt.Sprintf("%s|", strings.Repeat("-", size))
	}

	return separator
}
