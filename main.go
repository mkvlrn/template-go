// Package main is the application entrypoint.
package main

import (
	"fmt"
	"log"

	"github.com/mkvlrn/template-go/internal/math"
)

//nolint:mnd
func main() {
	fmt.Printf("%d + %d = %.2f\n", 1, 2, math.Add(1, 2))

	tenDividedByTwo, err := math.Divide(10, 2)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Printf("%d / %d = %.2f\n", 10, 2, tenDividedByTwo)
}
