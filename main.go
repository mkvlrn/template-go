// Package main is the application entrypoint.
package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/mkvlrn/template-go/internal/calculator"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "2 + 2", `The expression to be solved, e.g., 2 * 5`)
	flag.Parse()

	solution, err := calculator.Solve(input)
	if err != nil {
		panic(err)
	}

	formattedSolution := strconv.FormatFloat(solution, 'f', -1, 64)
	fmt.Printf("%s = %s\n", input, formattedSolution)
}
