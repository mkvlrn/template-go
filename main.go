// Package main is the entrypoint for the hello cmd.
package main

import (
	"fmt"
	"log"

	"github.com/mkvlrn/template-go/internal/math"
)

func main() {
	fmt.Printf("%d + %d = %f\n", 1, 2, math.Add(1, 2)) //nolint:mnd

	tenDividedByTwo, err := math.Divide(10, 2) //nolint:mnd
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Printf("%d / %d = %f\n", 10, 2, tenDividedByTwo) //nolint:mnd
}
