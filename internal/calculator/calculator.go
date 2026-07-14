// Package calculator parses input and returns the calculation done by the math package.
package calculator

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/mkvlrn/template-go/internal/simplemath"
)

type expression struct {
	a  float64
	op string
	b  float64
}

// Solve uses the math package to run the expected operation.
func Solve(input string) (float64, error) {
	var result float64

	xp, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	switch xp.op {
	case "+":
		result = simplemath.Add(xp.a, xp.b)

	case "-":
		result = simplemath.Subtract(xp.a, xp.b)

	case "*":
		result = simplemath.Multiply(xp.a, xp.b)

	case "/":
		result, err = simplemath.Divide(xp.a, xp.b)
		if err != nil {
			return 0, err
		}
	}

	return math.Round(result*100) / 100, nil //nolint:mnd
}

func parseInput(input string) (*expression, error) {
	regex := regexp.MustCompile(`([+-]?[0-9]+?[.]?[0-9]*?)\s([+/*-])\s([+-]?[0-9]+?[.]?[0-9]*)`)

	if !regex.MatchString(input) {
		return nil, fmt.Errorf("%q is an invalid expression", input)
	}

	parsed := regex.FindStringSubmatch(input)
	a, _ := strconv.ParseFloat(parsed[1], 64)
	b, _ := strconv.ParseFloat(parsed[3], 64)

	return &expression{a: a, b: b, op: parsed[2]}, nil
}
