// Package main is the entrypoint for the hello cmd.
package main

import (
	"fmt"

	"github.com/mkvlrn/template-go/internal/stuff"
)

func main() {
	fmt.Println(stuff.Hello())
}
